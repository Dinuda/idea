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
//Config to get DB properties
type Config struct{
	Username string
	Password string
	Protocol string
	URL string
	Port string
	Schema string
}

//func cognito(tokens)																																``

//Connect is used to connect to the db
func Connect(config Config)(*sql.DB, error){
	log.Println("Connecting to the DB")
	path := config.Username+":"+config.Password+"@"+config.Protocol+"("+config.URL+":"+config.Port+")/"+config.Schema
	db, err := sql.Open("mysql", path)
	err = db.Ping()
	if err != nil{
		return nil, err
	}
	log.Println("Connected to the DB")
	return db, err 
}

//Prepare is used to prepare the sql stmt
func Prepare()error{
	return nil
}

//AddUser adds a new user
func AddUser(user models.User)(int64, error){
	log.Println("Adding a new User to the DB")
	return 0, nil
}

//AddInvestor adds a new Investor
func AddInvestor(investor models.Investor)(int64, error){
	log.Println("Adding a new Investor User to the DB")
	return 0, nil
}

//AddStudent adds a new Investor
func AddStudent(student models.Student)(int64, error){
	log.Println("Adding a new Student User to the DB")
	return 0, nil
}