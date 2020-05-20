package service

import (
	"fmt"
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	"../models"
	"../repository"
)

type claims struct {
	ProjectID int 
	jwt.StandardClaims
}

var jwtKey = []byte(`gIQhHG6cxsMuyoR92KWvPmzUwd501BjY8fAZltJbE7aSeD4TXiNOCLVpnkrq3F`)

//GetProjectCategories gets all the Categories roles
func GetProjectCategories() (map[int]string, error) {
	projectCategories, err := repository.GetProjectCategories()
	if err != nil || projectCategories == nil {
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

	project.Host = userID
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

//GenarateProjectInvitationCode created a encrypted code to enable invitation code for the user
func GenarateProjectInvitationCode(projectID int)(string, error){
	exp := time.Now().Add(48 *time.Hour)
		claims := &claims{
			ProjectID: projectID,
			StandardClaims: jwt.StandardClaims{
				// In JWT, the expiry time is expressed as unix milliseconds
				ExpiresAt: exp.Unix(),
			},
		}
	
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
		tokenString, err := token.SignedString(jwtKey)

		if err != nil {
			log.Println("Error creating the token, " + err.Error())
			//http.Error(w, "Error creating the token, " + err.Error(), http.StatusInternalServerError)
			return "", fmt.Errorf("Error creating the token, " + err.Error())
		}
		return tokenString, nil
}

