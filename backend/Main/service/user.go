package service

import (
	"fmt"
	"log"

	"../models"
	"../pkg"
	"../repository"
)

//AddUser is used to and a new User
func AddUser(user models.User) error {
	AddInvestor := func(investor models.Investor) error {
		log.Println("Adding an Investor")
		rowsAffected, err := repository.AddInvestor(investor)
		if err != nil && rowsAffected < 1 {
			log.Println("ERR: adding new Investor failed")
			return err
		}
		return nil
	}
	AddStudent := func(student models.Student) error {
		log.Println("Adding an Student")
		rowsAffected, err := repository.AddStudent(student)
		if err != nil && rowsAffected < 1 {
			log.Println("ERR: adding new Investor failed")
			return err
		}
		return nil
	}

	log.Println("Adding a new User")
	hasedPass, err := pkg.GenarateHash(user.Password)
	user.Password = hasedPass
	if err != nil {
		log.Println("Error hashing the password")
		return err
	}

	userID, err := repository.AddUser(user)
	log.Println(userID)
	if err != nil && userID == 0 {
		log.Println("Error adding new User failed")
		return err
	}

	if user.Type == "Investor" {
		log.Println("New User is a Investor")
		user.Investor.UserID = userID
		err = AddInvestor(user.Investor)
		return err
	} else if user.Type == "Student" {
		log.Println("New User is a student")
		user.Student.UserID = userID
		err = AddStudent(user.Student)
		return err
	}

	return fmt.Errorf("User type not Given But a new user is added")

}

//GetUserType gets the user type investor|student
func GetUserType(user models.User) (string, error) {
	log.Println("getting user type")
	return "", nil
}

//GetProfessions gets all the Professions roles
func GetProfessions() ([]models.Profession, error) {
	log.Println("Getting Professions roles")
	professions, err := repository.GetProfessions()
	if err != nil || len(professions) < 1 {
		log.Println("Error retriving professions form the DB, " + err.Error())
		return []models.Profession{}, err
	}
	return professions, nil
}

//GetProjectCategories gets all the Categories roles
func GetProjectCategories() ([]models.ProjectCategory, error) {
	log.Println("Getting Category roles")
	ProjectCategories, err := repository.GetProjectCategories()
	if err != nil || len(ProjectCategories) < 1 {
		log.Println("Error retriving Categories form the DB, " + err.Error())
		return []models.ProjectCategory{}, err
	}
	return ProjectCategories, nil
}
