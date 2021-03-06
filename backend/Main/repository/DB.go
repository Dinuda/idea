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
	createProjectStudentTeamStmt,
	createProjectInvestorTeamStmt,
	insertStudentToProjectStudentTeamStmt,
	insertInvestorToProjectInvestorTeamStmt,
	selectUserStmt,
	selectUserIDStmt,
	selectInvestorStmt,
	selectPasswordStmt,
	selectProjectCatagoriesStmt,
	selectProjectStmt,
	selectProjectIDStmt,
	selectProjectStudentTeamStmt,
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
			name,
			email, 
			phone_no, 
			description, 
			type
		) 
		VALUES(?,?,?,?,?,?,?)`)
	if err != nil {
		return fmt.Errorf("Error preparing insertUserStmt, " + err.Error())
	}

	insertInvestorStmt, err = DB.Prepare(`INSERT INTO investors 
		(
			id,
			linkedin,
			company
		) 
		VALUES(?,?,?)`)
	if err != nil {
		return fmt.Errorf("Error preparing insertInvestorStmt, " + err.Error())
	}

	insertStudentStmt, err = DB.Prepare(`INSERT INTO students 
		(	
			id,
			profession, 
			cv
		) 
		VALUES(?,?,?)`)
	if err != nil {
		return fmt.Errorf("Error preparing insertStudentStmt, " + err.Error())
	}

	insertProjectStmt, err = DB.Prepare(`INSERT INTO projects 
		(
			title, 
			description,
			created_at,
			category,
			host
		)
		VALUES(?,?,NOW(),?,?)`)
	if err != nil {
		return fmt.Errorf("Error preparing insertProjectStmt, " + err.Error())
	}

	createProjectStudentTeamStmt, err = DB.Prepare(`INSERT INTO projectstudentteam
		(	
			project_id,
			user_id
		) 
		VALUES(?,?)`)
	if err != nil {
		return fmt.Errorf("Error preparing createStudentTeamStmt, " + err.Error())
	}

	insertStudentToProjectStudentTeamStmt, err = DB.Prepare(`INSERT INTO projectinvestorteam 
	(
		project_id, 
		user_id
	) 
	VALUES(?,?)`)
	if err != nil {
		return fmt.Errorf("Error preparing insertStudentToProjectStudentTeamTeamStmt, " + err.Error())
	}

	createProjectInvestorTeamStmt, err = DB.Prepare(`INSERT INTO projectinvestorteam
		(
			project_id,
			user_id
		) 
		VALUES(?,?)`)
	if err != nil {
		return fmt.Errorf("Error preparing createStudentTeamStmt, " + err.Error())
	}

	selectUserStmt, err = DB.Prepare(`SELECT
		name,
		email, 
		phone_no, 
		description, 
		type 
		FROM users WHERE username=?`)
	if err != nil {
		return fmt.Errorf("Error preparing selectUserStmt, " + err.Error())
	}

	selectUserIDStmt, err = DB.Prepare(`SELECT id FROM users WHERE username=?`)
	if err != nil {
		return fmt.Errorf("Error preparing selectUserIDStmt, " + err.Error())
	}

	selectInvestorStmt, err = DB.Prepare(`SELECT * FROM users INNER JOIN investors ON users.id=investors.id  WHERE users.id=?`)
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

	selectProjectStmt, err = DB.Prepare(`SELECT 
		projects.id,
		projects.title, 
		projects.description, 
		projects.created_at, 
		projects.category,
		users.id
		FROM users INNER JOIN projects ON users.id=projects.host WHERE users.username=?  `)
	if err != nil {
		return fmt.Errorf("Error preparing selectProjectStmt, " + err.Error())
	}

	selectProjectIDStmt, err = DB.Prepare(`SELECT 
		projects.id, 
		projects.host,
		users.id 
		FROM users INNER JOIN projects ON users.id=projects.host WHERE users.username=?`)
	if err != nil {
		return fmt.Errorf("Error preparing selectProjectIDStmt, " + err.Error())
	}
	// selectProjectStudentTeamStmt, err = DB.Prepare(`SELECT user_id FROM projectstudentteam INNER JOIN projects USING()`)
	// if err != nil {
	// 	return fmt.Errorf("Error preparing selectProjectStudentTeamStmt, " + err.Error())
	// }

	selectPasswordStmt, err = DB.Prepare(`SELECT password FROM users WHERE username = ?`)
	if err != nil {
		return fmt.Errorf("Error preparing the selectPasswordStmt, " + err.Error())
	}

	return nil
}
