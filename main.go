package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// User is the basic structure of a user record
type User struct {
	ID              string `json:"user_id"`
	UserName        string `json:"user_name"`
	UserEmail       string `json:"user_email"`
	FName           string `json:"user_first_name"`
	LName           string `json:"user_last_name"`
	Password        string `json:"password"`
	PasswordChange  string `json:"password_change"`
	PasswordExpired string `json:"password_expired"`
	LastLogon       string `json:"last_logon"`
	AccountLocked   string `json:"account_locked"`
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "cgdavis"
	dbPass := "DzftXvz$eR7VpY^h"
	dbServer := "tcp(172.17.232.252:3306)"
	dbName := "cms"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@"+dbServer+"/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

// GetUsers returns all records in the user table
func GetUsers(c *gin.Context) {
	db := dbConn()
	selDB, err := db.Query("CALL read_users()")
	if err != nil {
		panic(err.Error)
	}

	user := User{}
	users := []User{}
	for selDB.Next() {
		var id, username, useremail, fname, lname, password, passwordchange, passwordexpired, lastlogon, accountlocked string
		err = selDB.Scan(&id, &username, &useremail, &fname, &lname, &password, &passwordchange, &passwordexpired, &lastlogon, &accountlocked)
		if err != nil {
			log.Println(err)
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
		}
		user.ID = id
		user.UserName = username
		user.UserEmail = useremail
		user.FName = fname
		user.LName = lname
		user.Password = password
		user.PasswordChange = passwordchange
		user.PasswordExpired = passwordexpired
		user.LastLogon = lastlogon
		user.AccountLocked = accountlocked
		users = append(users, user)
	}

	c.JSON(200, gin.H{
		"result": users,
	})

	defer db.Close()
}

// GetUser returns a single user record from the database
func GetUser(c *gin.Context) {
	nID := c.Param("user_id")
	db := dbConn()
	selDB, err := db.Query("CALL read_user(?)", nID)
	if err != nil {
		panic(err.Error)
	}

	user := User{}
	users := []User{}
	for selDB.Next() {
		var id, username, useremail, fname, lname, password, passwordchange, passwordexpired, lastlogon, accountlocked string
		err = selDB.Scan(&id, &username, &useremail, &fname, &lname, &password, &passwordchange, &passwordexpired, &lastlogon, &accountlocked)
		if err != nil {
			log.Println(err)
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
		}
		user.ID = id
		user.UserName = username
		user.UserEmail = useremail
		user.FName = fname
		user.LName = lname
		user.Password = password
		user.PasswordChange = passwordchange
		user.PasswordExpired = passwordexpired
		user.LastLogon = lastlogon
		user.AccountLocked = accountlocked
		users = append(users, user)
	}

	c.JSON(200, gin.H{
		"result": users,
	})

	defer db.Close()
}

// GetUserUsername retrieves a record by username
func GetUserUsername(c *gin.Context) {
	nID := c.Param("user_name")
	db := dbConn()
	selDB, err := db.Query("CALL read_user_username(?)", nID)
	if err != nil {
		panic(err.Error)
	}

	user := User{}
	users := []User{}
	for selDB.Next() {
		var id, username, useremail, fname, lname, password, passwordchange, passwordexpired, lastlogon, accountlocked string
		err = selDB.Scan(&id, &username, &useremail, &fname, &lname, &password, &passwordchange, &passwordexpired, &lastlogon, &accountlocked)
		if err != nil {
			log.Println(err)
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
		}
		user.ID = id
		user.UserName = username
		user.UserEmail = useremail
		user.FName = fname
		user.LName = lname
		user.Password = password
		user.PasswordChange = passwordchange
		user.PasswordExpired = passwordexpired
		user.LastLogon = lastlogon
		user.AccountLocked = accountlocked
		users = append(users, user)
	}

	c.JSON(200, gin.H{
		"result": users,
	})

	defer db.Close()
}

// CreateUser adds user information to the database
func CreateUser(c *gin.Context) {
	db := dbConn()
	var user User
	if err := c.BindJSON(&user); err == nil {
		statement, _ := db.Prepare("CALL create_user (?, ?, ?, ?, ?, ?, ?, ?, ?)")
		statement.Exec(user.UserName, user.UserEmail, user.FName, user.LName, user.Password, user.PasswordChange, user.PasswordExpired, user.LastLogon, user.AccountLocked)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		}
	} else {
		c.String(http.StatusInternalServerError, err.Error())
	}
	defer db.Close()
}

// UpdateUser updates information in the user record
func UpdateUser(c *gin.Context) {
	db := dbConn()
	var user User
	if err := c.BindJSON(&user); err == nil {
		statement, _ := db.Prepare("CALL update_user (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
		statement.Exec(user.ID, user.UserName, user.UserEmail, user.FName, user.LName, user.Password, user.PasswordChange, user.PasswordExpired, user.LastLogon, user.AccountLocked)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		}
	} else {
		c.String(http.StatusInternalServerError, err.Error())
	}
	defer db.Close()
}

// DeleteUser deletes a user record from the database
func DeleteUser(c *gin.Context) {
	nID := c.Param("user_id")
	db := dbConn()
	statement, _ := db.Prepare("CALL delete_user(?)")
	statement.Exec(nID)
	defer db.Close()
}

func main() {
	r := gin.Default()

	r.GET("/v1/getusers", GetUsers)
	r.GET("/v1/getuser/:user_id", GetUser)
	r.GET("/v1/getuserusername/:user_name", GetUserUsername)
	r.GET("/v1/deleteuser/:user_id", DeleteUser)
	r.POST("/v1/createuser", CreateUser)
	r.POST("/v1/updateuser", UpdateUser)

	r.Run(":8000")
}
