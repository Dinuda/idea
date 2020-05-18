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

	project.HostID = userID
	var lastInsertID int
	lastInsertID, err = repository.CreateProject(project)
	if err != nil{
		log.Println("Error creating a project, " + err.Error())
		return models.Project{}, fmt.Errorf("Error creating a project, " + err.Error())
	}

	err = repository.CreateProjectStudentTeam(lastInsertID, userID)
	if err != nil {
		log.Println("Error creating a projectstudentteam," + err.Error())
		return models.Project{}, fmt.Errorf("Error creating a studentteam, " + err.Error())
	}
	return project, nil

}

//AddStudentToProjectStudentTeam adds a new student to the studentteam of the project
func AddStudentToProjectStudentTeam(username string, projectID int) (error){
	userID, err := repository.GetUserID(username)
	if err != nil {
		log.Println("Error getting the user id, " + err.Error())
		return fmt.Errorf("Error getting the user id, " + err.Error())
	}
	rowsAffected, err := repository.AddStudentToProjectStudentTeam(projectID, userID)
	if err != nil || rowsAffected  < 1{
		return err
	}
	return nil
}

//GetProjects gets all the project of the student
func GetProjects(username string)([]models.Project, error){
	projects, err := repository.GetProjects(username)
	if err != nil{
		log.Println("Error getting projects, " + err.Error())
		return projects, err
	}
	return projects, nil
}