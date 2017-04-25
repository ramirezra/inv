package config

import (
	"database/sql"
	"fmt"
)

// DB exported
var DB *sql.DB

var driver = "postgres"
var username = "bond"
var pword = "password"
var datastore = "localhost/bookstore"

func init() {
	var err error
	// DB, err = sql.Open("postgres", "postgres://bond:password@localhost/bookstore?sslmode=disable")
	DB, err = sql.Open(driver, driver+"://"+username+":"+pword+"@"+datastore+"?sslmode=disable")
	if err != nil {
		panic(err)
	}
	if err = DB.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("You connected to your database.")

}
