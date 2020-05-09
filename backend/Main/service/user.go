package service

import (
	"log"
	"fmt"

	"../models"
	"../repository"
)

//AddUser is used to and a new User
func AddUser(user models.User)error{
	AddInvestor := func(investor models.Investor)error{
		log.Println("Adding an Investor")
		rowsAffected, err := repository.AddInvestor(investor)
		if err != nil && rowsAffected < 1 {
			log.Println("ERR: adding new Investor failed")
			return err
		}
		return nil
	}
	AddStudent := func(student models.Student)error{
		log.Println("Adding an Student")
		rowsAffected, err := repository.AddStudent(student)
		if err != nil && rowsAffected < 1 {
			log.Println("ERR: adding new Investor failed")
			return err
		}
		return nil
	}


	log.Println("Adding a new User")
	rowsAffected, err := repository.AddUser(user)
	if err != nil && rowsAffected < 1 {
		log.Println("ERR: adding new User failed")
		return err
	}

	if user.Type == "Investor"{
		log.Println("New User is a Investor")
		err = AddInvestor(user.Investor)
		return err
	}else if user.Type == "Student"{
		log.Println("New User is a student")
		err = AddStudent(user.Student)
		return err
	}

	return fmt.Errorf("User type not Given But a new user is added")

	
}

//AddInvestor adds a new Investor
