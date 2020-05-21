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
		err = addInvestor(user.Investor)
		return err
	} else if user.Type == "Student" {
		log.Println("New User is a student")
		user.Student.UserID = userID
		err = addStudent(user.Student)
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
	professions, err := repository.GetProfessions()
	if err != nil || len(professions) < 1 {
		log.Println("Error retriving professions form the DB, " + err.Error())
		return professions, err
	}
	return professions, nil
}


//GetUser gets all the data from the user
func GetUser(username string)(models.User, error){
	user, err := repository.GetUser(username)
	if err != nil {
		log.Println("Error getting user,", username, " error ", err.Error())
		return user, err
	}
	return user, nil
}