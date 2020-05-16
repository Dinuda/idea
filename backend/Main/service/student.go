package service

import (
	"log"

	"../repository"
	"../models"
)

func addStudent(student models.Student) error {
	log.Println("Adding an Student")
	rowsAffected, err := repository.AddStudent(student)
	if err != nil && rowsAffected < 1 {
		log.Println("ERR: adding new Investor failed")
		return err
	}
	return nil
}