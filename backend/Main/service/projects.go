package service

import (
	"fmt"
	"log"

	"../models"
	"../repository"
)

//GetProjectCategories gets all the Categories roles
func GetProjectCategories() ([]models.ProjectCategory, error) {
	projectCategories, err := repository.GetProjectCategories()
	if err != nil || len(projectCategories) < 1 {
		log.Println("Error retriving Categories form the DB, " + err.Error())
		return projectCategories, err
	}
	return projectCategories, nil
}

//CreateProject creates new projects with the teams it requires
func CreateProject(project models.Project, username string) (models.Project, error) {
	userID, err := repository.GetUserID(username)
	if err != nil {
		log.Println("Error getting the user id, " + err.Error())
		return models.Project{}, fmt.Errorf("Error getting the user id, " + err.Error())
	}

	var lastInsertID , rowsAffected int
	lastInsertID, err = repository.CreateStudentTeam(userID)
	if err != nil {
		log.Println("Error creating a studentteam," + err.Error())
		return models.Project{}, fmt.Errorf("Error creating a studentteam, " + err.Error())
	}
	project.StudentTeamID = lastInsertID

	rowsAffected, err = repository.CreateProject(project)
	if err != nil || rowsAffected < 1{
		log.Println("Error creating a project, " + err.Error())
		return models.Project{}, fmt.Errorf("Error creating a project, " + err.Error())
	}
	return project, nil

}

//AddStudentToTeam adds a new student to the studentteam of the project
func AddStudentToTeam(username string, teamID int) (error){
	userID, err := repository.GetUserID(username)
	if err != nil {
		log.Println("Error getting the user id, " + err.Error())
		return fmt.Errorf("Error getting the user id, " + err.Error())
	}
	rowsAffected, err := repository.AddStudentToTeam(teamID, userID)
	if err != nil || rowsAffected  < 1{
		return err
	}
	return nil
}