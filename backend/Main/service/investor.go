package service

import (
	"log"

	"../repository"
	"../models"
)

func addInvestor(investor models.Investor) error {
	log.Println("Adding an Investor")
	rowsAffected, err := repository.AddInvestor(investor)
	if err != nil && rowsAffected < 1 {
		log.Println("ERR: adding new Investor failed")
		return err
	}
	return nil
}
