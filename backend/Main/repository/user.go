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

var (
	insertUserStmt,
	insertInvestorStmt,
	insertStudentStmt *sql.Stmt
)

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
	var err error
	insertUserStmt, err = DB.Prepare(`INSERT INTO users 
		(
			username, 
			password, 
			firstname, 
			lastname, 
			email, 
			phone_no, 
			date_of_birth, 
			description, 
			type
		) 
		VALUES(?,?,?,?,?,?,?,?,?)`)
	if err != nil{
		return err
	}
	insertInvestorStmt, err = DB.Prepare(`INSERT INTO users 
		(
			linkedin, 
			company
		) 
		VALUES(?,?,?,?)`)
	if err != nil{
		return err
	}
	insertStudentStmt, err = DB.Prepare(`INSERT INTO users 
		(
			profession, 
			university, 
			cv,
			team_role 
		) 
		VALUES(?,?,?,?)`)
	if err != nil{
		return err
	}
	return nil
}

//AddUser adds a new user
func AddUser(user models.User)(int64, error){
	log.Println("Adding a new User to the DB")
	result, err := insertUserStmt.Exec(
		user.Username,
		user.Password,
		user.Firstname,
		user.Lastname,
		user.Email,
		user.PhoneNo,
		user.DateofBirth,
		user.Description,
		user.Type,
	)
	rowsAffected, _ := result.RowsAffected()
	return rowsAffected(), err
}

//AddInvestor adds a new Investor
func AddInvestor(investor models.Investor)(int64, error){
	log.Println("Adding a new Investor User to the DB")
	result, err := insertInvestorStmt.Exec(
		investor.Linkedin,
		investor.Company,
	)
	rowsAffected, _ := result.RowsAffected()
	return rowsAffected, err
}

//AddStudent adds a new Investor
func AddStudent(student models.Student)(int64, error){
	log.Println("Adding a new Student User to the DB")
	result, err := insertStudentStmt.Exec(
		student.Profession,
		student.University,
		student.cv,
		student.TeamRole
	)
	return 0, nil
}