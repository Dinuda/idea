package repository

import(
	"log"
	"database/sql"

	//Driver to intrigate with MySql
	_ "github.com/go-sql-driver/mysql"

	"../models"
)
//DB DB connection
var DB *sql.DB

//Connect is used to connect to the db
func Connect(path string)(*sql.DB, error){
	log.Println("INFO: Connecting to the DB")
	db, err := sql.Open("mysql", path)
	log.Println("INFO: Connected to the DB")
	return db, err 
}

//Prepare is used to prepare the sql stmt
func Prepare()error{
	return nil
}

//AddEntrepreneur adds a new entrepreneur
func AddEntrepreneur(User models.Entrepreneur)(int64, error){
	log.Println("Adding a new Entrepreneur")
	return 0, nil
}