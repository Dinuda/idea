package repository

import (
	"database/sql"
	"fmt"
	"log"

	//Driver to intrigate with MySql
	_ "github.com/go-sql-driver/mysql"
)

//DB DB connection
var DB *sql.DB

//DNS to get DB properties
var DNS string

//prepare stmt
var (
	insertUserStmt,
	insertInvestorStmt,
	insertStudentStmt,
	insertProjectStmt,
	createStudentTeamStmt,
	createInvestorTeamStmt,
	insertStudentToTeamStmt,
	selectUserStmt,
	selectUserIDStmt,
	selectInvestorStmt,
	selectPasswordStmt,
	selectProjectCatagoriesStmt,
	selectProfessionsStmt *sql.Stmt
)

//Connect is used to connect to the db
func Connect() (*sql.DB, error) {
	log.Println("Connecting to the DB")
	db, err := sql.Open("mysql", DNS)
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	log.Println("Connected to the DB")
	return db, err
}

//Prepare is used to prepare the sql stmt
func Prepare() error {
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
	if err != nil {
		return fmt.Errorf("Error preparing insertUserStmt, " + err.Error())
	}

	insertInvestorStmt, err = DB.Prepare(`INSERT INTO investors 
		(
			user_id,
			linkedin, 
			company
		) 
		VALUES(?,?,?)`)
	if err != nil {
		return fmt.Errorf("Error preparing insertInvestorStmt, " + err.Error())
	}

	insertStudentStmt, err = DB.Prepare(`INSERT INTO students 
		(
			profession, 
			cv
		) 
		VALUES(?,?)`)
	if err != nil {
		return fmt.Errorf("Error preparing insertStudentStmt, " + err.Error())
	}

	insertProjectStmt, err = DB.Prepare(`INSERT INTO projects 
		(
			title, 
			description,
			created_date,
			category,
			studentteam_id
		)
		VALUES(?,?,?,?,?)`)
	if err != nil {
		return fmt.Errorf("Error preparing insertProjectStmt, " + err.Error())
	}

	createStudentTeamStmt, err = DB.Prepare(`INSERT INTO studentteam
		(
			student_id
		) 
		VALUES(?)`)
	if err != nil {
		return fmt.Errorf("Error preparing createStudentTeamStmt, " + err.Error())
	}

	insertStudentToTeamStmt, err = DB.Prepare(`INSERT INTO studentteam (
		id, 
		user_id
		) VALUES(?,?)`)
	if err != nil {
		return fmt.Errorf("Error preparing insertStudentToTeamStmt, " + err.Error())
	}

	createInvestorTeamStmt, err = DB.Prepare(`INSERT INTO investorteam
		(
			investor_id
		) 
		VALUES(?)`)
	if err != nil {
		return fmt.Errorf("Error preparing createStudentTeamStmt, " + err.Error())
	}

	selectUserStmt, err = DB.Prepare(`SELECT
		firstname, 
		lastname, 
		email, 
		phone_no, 
		date_of_birth, 
		description, 
		type 
		FROM users WHERE id=?`)
	if err != nil {
		return fmt.Errorf("Error preparing selectUserStmt, " + err.Error())
	}

	selectUserIDStmt, err = DB.Prepare(`SELECT id FROM users WHERE username=?`)
	if err != nil {
		return fmt.Errorf("Error preparing selectUserIDStmt, " + err.Error())
	}

	selectInvestorStmt, err = DB.Prepare(`SELECT * FROM users INNER JOIN investors ON users.id=investors.user_id  WHERE users.id=?`)
	if err != nil {
		return fmt.Errorf("Error preparing selectInvestorStmt, " + err.Error())
	}

	selectProfessionsStmt, err = DB.Prepare(`SELECT * FROM professions`)
	if err != nil {
		return fmt.Errorf("Error preparing selectProfessionsStmt, " + err.Error())
	}

	selectProjectCatagoriesStmt, err = DB.Prepare(`SELECT * FROM categories`)
	if err != nil {
		return fmt.Errorf("Error preparing selectProjectCategoryStmt, " + err.Error())
	}

	selectPasswordStmt, err = DB.Prepare(`SELECT password FROM users WHERE username = ?`)
	if err != nil {
		return fmt.Errorf("Error preparing the selectPasswordStmt, " + err.Error())
	}

	return nil
}
