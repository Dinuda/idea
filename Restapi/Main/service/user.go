package service

import (
	"log"

	"../models"
	"../repository"
)

//AddEntrepreneur adds a new entrepreneur
func AddEntrepreneur(User models.Entrepreneur)error{
	log.Println("Adding an Entrepreneur")
	rowsAffected, err := repository.AddEntrepreneur(User)
	if err != nil && rowsAffected < 1 {
		log.Println("ERR: adding new entrepreneur failed")
		return err
	}
	return nil
}