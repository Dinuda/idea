package repository

import(
	"log"
	"fmt"
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
	insertStudentStmt,
	selectUserStmt,
	selectProfessionsStmt *sql.Stmt
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
		VALUES(?,?,?,?,?,?,?,?,?)
		SELECT LAST_INSERT_ID`)
	if err != nil{
		return fmt.Errorf("Error preparing insertUserStmt, "+ err.Error())
	}
	insertInvestorStmt, err = DB.Prepare(`INSERT INTO investors 
		(
			user_id
			linkedin, 
			company
		) 
		VALUES(?,?,?)`)
	if err != nil{
		return fmt.Errorf("Error preparing insertInvestorStmt, "+ err.Error())
	}
	insertStudentStmt, err = DB.Prepare(`INSERT INTO students 
		(
			profession, 
			university, 
			cv,
			team_role 
		) 
		VALUES(?,?,?,?)`)
	if err != nil{
		return fmt.Errorf("Error preparing insertStudentStmt, "+ err.Error())
	}

	selectUserStmt, err = DB.Prepare(`SELECT * FROM users WHERE type=investor INNER JOIN investors ON users.id=investors.user_id`)
	if err != nil{
		return fmt.Errorf("Error preparing selectUserStmt, "+ err.Error())
	}
	selectProfessionsStmt, err = DB.Prepare(`SELECT name FROM professions`)
	if err != nil{
		return fmt.Errorf("Error preparing selectProfessionsStmt, "+ err.Error())
	
	}
	return nil
}

//AddUser adds a new user
func AddUser(user models.User)(int, error){
	log.Println("Adding a new User to the DB")
	result, err := insertUserStmt.Query(
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
	var UserID int
	err = result.Scan(&UserID)
	return UserID, err
}

//AddInvestor adds a new Investor
func AddInvestor(investor models.Investor)(int64, error){
	log.Println("Adding a new Investor User to the DB")
	result, err := insertInvestorStmt.Exec(
		investor.UserID,
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
		student.CV,
		student.TeamRole,
	)
	rowsAffected, _ := result.RowsAffected()
	return rowsAffected, err
}

//GetProfessions get all the professions from the db
// func GetProfessions()([]string, error){
// 	result, err := selectProfessions.Query()
// 	var professions []string
// 	for result.Next(){
// 		var profession string
// 		result.Scan(&profession)
// 		professions = professions.addend(profession, professions)
// 	}
// }