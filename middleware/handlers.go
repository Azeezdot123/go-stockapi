package middleware

import "database/sql"

type response struct{
	ID			int64	`json:"id,omitempty"`
	Message		string	`json:"message,omitempty"` 
}

func CreateConnection() *sql.Db{
	err := getdotenv.Load(".env")
}
func GetStock(){}
func GetAllStock(){}
func CreateStock(){}
func UpdateStock(){}
func DeleteStock(){}