package service

import (
	"log"

	"../models"
	"../repository"
)

//AddInvestor adds a new Investor
func AddInvestor(User models.Investor)error{
	log.Println("Adding an Investor")
	rowsAffected, err := repository.AddInvestor(User)
	if err != nil && rowsAffected < 1 {
		log.Println("ERR: adding new Investor failed")
		return err
	}
	return nil
}