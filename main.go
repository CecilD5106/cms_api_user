package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// Person is the basic structure of a person record
type Person struct {
	ID    string `json:"person_id"`
	FName string `json:"first_name"`
	LName string `json:"last_name"`
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "cgdavis"
	dbPass := "DzftXvz$eR7VpY^h"
	dbServer := "tcp(172.17.232.252:3306)"
	dbName := "people"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@"+dbServer+"/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

// GetPeople returns all records in the person table
func GetPeople(c *gin.Context) {
	db := dbConn()
	selDB, err := db.Query("CALL read_all_people()")
	if err != nil {
		panic(err.Error)
	}

	person := Person{}
	people := []Person{}
	for selDB.Next() {
		var id, fname, lname string
		err = selDB.Scan(&id, &fname, &lname)
		if err != nil {
			log.Println(err)
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
		}
		person.ID = id
		person.FName = fname
		person.LName = lname
		people = append(people, person)
	}

	c.JSON(200, gin.H{
		"result": people,
	})

	defer db.Close()
}

// GetPerson returns a single person record from the database
func GetPerson(c *gin.Context) {
	nID := c.Param("person_id")
	db := dbConn()
	selDB, err := db.Query("CALL read_person(?)", nID)
	if err != nil {
		panic(err.Error)
	}

	indiv := Person{}
	people := []Person{}
	for selDB.Next() {
		var id, fname, lname string
		err = selDB.Scan(&id, &fname, &lname)
		if err != nil {
			log.Println(err)
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
		}
		indiv.ID = id
		indiv.FName = fname
		indiv.LName = lname
		people = append(people, indiv)
	}

	c.JSON(200, gin.H{
		"result": people,
	})

	defer db.Close()
}

// CreatePerson adds person information to the database
func CreatePerson(c *gin.Context) {
	db := dbConn()
	var person Person
	if err := c.BindJSON(&person); err == nil {
		statement, _ := db.Prepare("CALL create_person (?, ?)")
		statement.Exec(person.FName, person.LName)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		}
	} else {
		c.String(http.StatusInternalServerError, err.Error())
	}
	defer db.Close()
}

// UpdatePerson updates information in the person record
func UpdatePerson(c *gin.Context) {
	db := dbConn()
	var person Person
	if err := c.BindJSON(&person); err == nil {
		statement, _ := db.Prepare("CALL update_person (?, ?, ?)")
		statement.Exec(person.ID, person.FName, person.LName)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		}
	} else {
		c.String(http.StatusInternalServerError, err.Error())
	}
	defer db.Close()
}

// DeletePerson deletes a person record from the database
func DeletePerson(c *gin.Context) {
	nID := c.Param("person_id")
	db := dbConn()
	statement, _ := db.Prepare("CALL delete_person(?)")
	statement.Exec(nID)
	defer db.Close()
}

func main() {
	r := gin.Default()

	r.GET("/getpeople", GetPeople)
	r.GET("/getperson/:person_id", GetPerson)
	r.GET("/deleteperson/:person_id", DeletePerson)
	r.POST("/createperson", CreatePerson)
	r.POST("/updateperson", UpdatePerson)

	r.Run(":8000")
}
