package handlers

import (
	"../config"
	"fmt"
	"database/sql"
)

func InitDb() {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", config.DB_USER, config.DB_PASSWORD, config.DB_NAME, config.DB_HOST, config.PORT)
	var err error
	db, err = sql.Open("postgres", dbinfo)
	checkErr(err)
	err = db.Ping()
	checkErr(err)
	fmt.Println("Successfully connected!")
}

func checkErr(err error)  {
	if err != nil {
		panic(err)
	}
}
