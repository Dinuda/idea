package repository

import(
	"log"
	"database/sql"

	//Driver to intrigate with MySql
	_ "github.com/go-sql-driver/mysql"
	//cognito "github.com/aws/aws-sdk-go/service/cognitoidentity"

	"../models"
)
//DB DB connection
var DB *sql.DB

//func cognito(tokens)																																``

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

//AddInvestor adds a new Investor
func AddInvestor(User models.Investor)(int64, error){
	log.Println("Adding a new Investor hello")
	return 0, nil
}