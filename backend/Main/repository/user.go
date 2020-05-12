package repository

import (
	"database/sql"
	"fmt"
	"log"
	//Driver to intrigate with MySql
	_ "github.com/go-sql-driver/mysql"

	"../models"
)

//DB DB connection
var DB *sql.DB

//DNS to get DB properties
var DNS string

var (
	insertUserStmt,
	insertInvestorStmt,
	insertStudentStmt,
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
			university, 
			cv,
			team_role 
		) 
		VALUES(?,?,?,?)`)
	if err != nil {
		return fmt.Errorf("Error preparing insertStudentStmt, " + err.Error())
	}

	selectInvestorStmt, err = DB.Prepare(`SELECT * FROM users INNER JOIN investors ON users.id=investors.user_id  WHERE users.id=?`)
	if err != nil {
		return fmt.Errorf("Error preparing selectUserStmt, " + err.Error())
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

//AddUser adds a new user
func AddUser(user models.User) (int, error) {
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
	if err != nil {
		return 0, err
	}
	UserID, errID := result.LastInsertId()
	if errID != nil {
		return 0, errID
	}
	return int(UserID), nil
}

//AddInvestor adds a new Investor
func AddInvestor(investor models.Investor) (int64, error) {
	log.Println("Adding a new Investor User to the DB")
	result, err := insertInvestorStmt.Exec(
		investor.UserID,
		investor.Linkedin,
		investor.Company,
	)
	if err != nil {
		return 0, err
	}
	rowsAffected, errResult := result.RowsAffected()
	if errResult != nil {
		return 0, err
	}
	return rowsAffected, nil
}

//AddStudent adds a new Investor
func AddStudent(student models.Student) (int64, error) {
	log.Println("Adding a new Student User to the DB")
	result, err := insertStudentStmt.Exec(
		student.Profession,
		student.University,
		student.CV,
		student.TeamRole,
	)
	if err != nil {
		return 0, err
	}
	rowsAffected, errResult := result.RowsAffected()
	if errResult != nil {
		return 0, err
	}
	return rowsAffected, nil
}

//GetProfessions get all the professions from the db
func GetProfessions() ([]models.Profession, error) {
	result, err := selectProfessionsStmt.Query()
	if err != nil {
		return []models.Profession{}, err
	}
	var professions []models.Profession
	for result.Next() {
		var profession models.Profession
		err = result.Scan(&profession.ID, &profession.Name)
		if err != nil{
			return []models.Profession{}, err
		}
		professions = append(professions, profession)
	}
	
	return professions, err
}

//GetProjectCategories get all the category from the db
func GetProjectCategories() ([]models.ProjectCategory, error) {
	result, err := selectProjectCatagoriesStmt.Query()
	if err != nil {
		return []models.ProjectCategory{}, err
	}
	var projectCategories []models.ProjectCategory
	for result.Next() {
		var projectCategory models.ProjectCategory
		err = result.Scan(&projectCategory.ID, &projectCategory.Name)
		if err != nil{
			return []models.ProjectCategory{}, err
		}
		projectCategories = append(projectCategories, projectCategory)
	}
	
	return projectCategories, err
}

//GetUserPassword retrieves password for a specific user
func GetUserPassword(username string) (string, error) {
	var password string
	err := selectPasswordStmt.QueryRow(username).Scan(&password)
	return password, err
}
