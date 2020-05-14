package service


import (
	"log"

	"../models"
	"../repository"
)

//GetProjectCategories gets all the Categories roles
func GetProjectCategories() ([]models.ProjectCategory, error) {
	log.Println("Getting Category roles")
	projectCategories, err := repository.GetProjectCategories()
	if err != nil || len(projectCategories) < 1 {
		log.Println("Error retriving Categories form the DB, " + err.Error())
		return projectCategories, err
	}
	return projectCategories, nil
}