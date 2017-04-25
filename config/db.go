package config

import (
	"database/sql"
	"fmt"
)

// DB exported
var DB *sql.DB

func init() {
	var err error

	DB, err = sql.Open(Driver, Driver+"://"+Account+":"+Pword+"@"+Datastore+"?sslmode=disable")
	if err != nil {
		panic(err)
	}
	if err = DB.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("You connected to your database.")
}
