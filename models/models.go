package models

import (
	"database/sql"
	"fmt"
	// "log"
	// "os"
	// "github.com/joho/godotenv"
)


type Stock struct{
	StockId	int64	`json:"stockid"`
	Name	string	`json:"name"`
	Price	int64	`json:"price"`
	Company	string	`json:"company"`	
}

func CreateConnection() *sql.DB {
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	// this is just using for testing it is not a production code .
	POSTGREURL := "postgres://gitpod:gitpod@localhost:5432/stockdb"

	db, err := sql.Open("postgres", POSTGREURL)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Succesfully Connected to Database")
	return db
}