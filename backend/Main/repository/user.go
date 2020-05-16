package repository

import (
	"log"

	"../models"
)

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
		student.UserID,
		student.Profession,
		student.CV,
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
	var professions []models.Profession
	result, err := selectProfessionsStmt.Query()
	if err != nil {
		return professions, err
	}
	
	for result.Next() {
		var profession models.Profession
		err = result.Scan(&profession.ID, &profession.Name)
		if err != nil {
			return []models.Profession{}, err
		}
		professions = append(professions, profession)
	}

	return professions, err
}

//GetUser get all the infomation of the User
func GetUser(username string)(models.User, error){
	var user models.User
	err := selectUserStmt.QueryRow(username).Scan(
		&user.Firstname,
		&user.Lastname,
		&user.Email,
		&user.PhoneNo,
		&user.Description,
		&user.Type,
	)
	if err != nil {
		return user, err
	}
	return user, nil
	
}


//GetUserPassword retrieves password for a specific user
func GetUserPassword(username string) (string, error) {
	var password string
	err := selectPasswordStmt.QueryRow(username).Scan(&password)
	return password, err
}

//GetUserID gets the user id
func GetUserID(username string) (int, error){
	var userID int
	err := selectUserIDStmt.QueryRow(username).Scan(&userID)
	return userID, err
}
