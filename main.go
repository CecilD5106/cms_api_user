package main

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

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
	AccessList      []Access
}

//Access is the level of access a user has
type Access struct {
	AccessID         string `json:"access_id"`
	IDUser           string `json:"user_id"`
	IDCourt          string `json:"court_id"`
	CaseAccess       string `json:"case_access"`
	PersonAccess     string `json:"person_access"`
	AccountingAccess string `json:"accounting_access"`
	JuryAccess       string `json:"jury_access"`
	AttorneyAccess   string `json:"attorney_access"`
	ConfigAccess     string `json:"configuration_access"`
	SecurityLevel    string `json:"security_level"`
	SealedCase       string `json:"sealed_case"`
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "masterAdmin"
	dbPass := "8JqpbWNsJ"
	dbServer := "tcp(cms-mysql-5106.cd0zye2bcjt9.us-west-2.rds.amazonaws.com:3306)"
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
		iid, err := strconv.Atoi(id)
		if err != nil {
			panic(err.Error)
		}
		selDB02, err := db.Query("CALL read_access_userid(?)", iid)
		if err != nil {
			panic(err.Error)
		}
		access := Access{}
		accessList := []Access{}
		for selDB02.Next() {
			var accessid, userid, courtid, caseaccess, personaccess, accountingaccess, juryaccess, attorneyaccess, configaccess, securitylevel, sealedcase string
			err := selDB02.Scan(&accessid, &userid, &courtid, &caseaccess, &personaccess, &accountingaccess, &juryaccess, &attorneyaccess, &configaccess, &securitylevel, &sealedcase)
			if err != nil {
				log.Println(err)
				c.JSON(500, gin.H{
					"error": err.Error(),
				})
			}
			access.AccessID = accessid
			access.IDUser = userid
			access.IDCourt = courtid
			access.CaseAccess = caseaccess
			access.PersonAccess = personaccess
			access.AccountingAccess = accountingaccess
			access.JuryAccess = juryaccess
			access.AttorneyAccess = attorneyaccess
			access.ConfigAccess = configaccess
			access.SecurityLevel = securitylevel
			access.SealedCase = sealedcase
			accessList = append(accessList, access)
		}
		user.AccessList = accessList
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
		iid, err := strconv.Atoi(id)
		if err != nil {
			panic(err.Error)
		}
		selDB02, err := db.Query("CALL read_access_userid(?)", iid)
		if err != nil {
			panic(err.Error)
		}
		access := Access{}
		accessList := []Access{}
		for selDB02.Next() {
			var accessid, userid, courtid, caseaccess, personaccess, accountingaccess, juryaccess, attorneyaccess, configaccess, securitylevel, sealedcase string
			err := selDB02.Scan(&accessid, &userid, &courtid, &caseaccess, &personaccess, &accountingaccess, &juryaccess, &attorneyaccess, &configaccess, &securitylevel, &sealedcase)
			if err != nil {
				log.Println(err)
				c.JSON(500, gin.H{
					"error": err.Error(),
				})
			}
			access.AccessID = accessid
			access.IDUser = userid
			access.IDCourt = courtid
			access.CaseAccess = caseaccess
			access.PersonAccess = personaccess
			access.AccountingAccess = accountingaccess
			access.JuryAccess = juryaccess
			access.AttorneyAccess = attorneyaccess
			access.ConfigAccess = configaccess
			access.SecurityLevel = securitylevel
			access.SealedCase = sealedcase
			accessList = append(accessList, access)
		}
		user.AccessList = accessList
		users = append(users, user)
	}

	c.JSON(200, gin.H{
		"result": users,
	})

	defer db.Close()
}

// GetUserUsername retrieves a record by username
func GetUserUsername(c *gin.Context) {
	userName := c.Param("user_name")
	db := dbConn()
	selDB, err := db.Query("CALL read_user_username(?)", userName)
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
		iid, err := strconv.Atoi(id)
		if err != nil {
			panic(err.Error)
		}
		selDB02, err := db.Query("CALL read_access_userid(?)", iid)
		if err != nil {
			panic(err.Error)
		}
		access := Access{}
		accessList := []Access{}
		for selDB02.Next() {
			var accessid, userid, courtid, caseaccess, personaccess, accountingaccess, juryaccess, attorneyaccess, configaccess, securitylevel, sealedcase string
			err := selDB02.Scan(&accessid, &userid, &courtid, &caseaccess, &personaccess, &accountingaccess, &juryaccess, &attorneyaccess, &configaccess, &securitylevel, &sealedcase)
			if err != nil {
				log.Println(err)
				c.JSON(500, gin.H{
					"error": err.Error(),
				})
			}
			access.AccessID = accessid
			access.IDUser = userid
			access.IDCourt = courtid
			access.CaseAccess = caseaccess
			access.PersonAccess = personaccess
			access.AccountingAccess = accountingaccess
			access.JuryAccess = juryaccess
			access.AttorneyAccess = attorneyaccess
			access.ConfigAccess = configaccess
			access.SecurityLevel = securitylevel
			access.SealedCase = sealedcase
			accessList = append(accessList, access)
		}
		user.AccessList = accessList
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
