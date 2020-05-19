package repository

import (
	"fmt"

	"../models"
)

//GetProjectCategories get all the category from the db
func GetProjectCategories() ([]models.ProjectCategory, error) {
	result, err := selectProjectCatagoriesStmt.Query()
	if err != nil {

		return []models.ProjectCategory{}, fmt.Errorf("Error getting ProjectCategories, " + err.Error())
	}
	var projectCategories []models.ProjectCategory
	for result.Next() {
		var projectCategory models.ProjectCategory
		err = result.Scan(&projectCategory.ID, &projectCategory.Name)
		if err != nil {
			return []models.ProjectCategory{}, fmt.Errorf("Error getting ProjectCategories, " + err.Error())
		}
		projectCategories = append(projectCategories, projectCategory)
	}

	return projectCategories, nil
}

//CreateProjectStudentTeam creates a new student team for the projects
func CreateProjectStudentTeam(projectID int, userID int) (error) {
	_, err := createProjectStudentTeamStmt.Exec(projectID, userID)
	if err != nil {
		return fmt.Errorf("Error creating projectstudentteam, " + err.Error())
	}
	return  nil
}

//CreateProject creates a new Project
func CreateProject(project models.Project) (int, error) {
	result, err := insertProjectStmt.Exec(
		project.Title,
		project.Description,
		project.Category,
		project.Host,
	)
	if err != nil {
		return 0, fmt.Errorf("Error creating a project, " + err.Error())
	}
	var lastInsertID int64
	lastInsertID, err = result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("Error getting no of LastInsertID," + err.Error())
	}
	return int(lastInsertID), nil
}

//AddStudentToProjectStudentTeam adds a new student to the studentteam of the project
func AddStudentToProjectStudentTeam(teamID int, userID int) (int, error) {
	result, err := insertStudentToProjectStudentTeamStmt.Exec(teamID, userID)
	if err != nil {
		return 0, fmt.Errorf("Error while adding student to the team, " + err.Error())
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("Error while getting the affected rows, " + err.Error())
	}
	return int(rowsAffected), nil
}

//GetProjects Gets all the project of the host
func GetProjects(username string)([]models.Project, error){
	var projects []models.Project
	result, err := selectProjectStmt.Query(username)
	if err != nil {
		return projects, err
	}
	
	for result.Next() {
		var project models.Project
		err = result.Scan(
			&project.ID,
			&project.Title,
			&project.Description,
			&project.CreatedDate,
			&project.Category,
			&project.Host,
			)
		if err != nil {
			return projects, err
		}
		projects = append(projects, project)
	}

	return projects, err
}


//GetStudentTeamID returns the studentteamid of that user
func GetStudentTeamID(hostID int) (int, error){
	return 0, nil
}

//GetProjectID get all the professions from the db
func GetProjectID(username string) ([]models.Project, error) {
	var projects []models.Project
	result, err := selectProjectIDStmt.Query(username)
	if err != nil {
		return projects, err
	}
	
	for result.Next() {
		var project models.Project
		var userID int
		err = result.Scan(&project.ID, &project.Host, &userID)
		if err != nil {
			return []models.Project{}, err
		}
		projects = append(projects, project)
	}

	return projects, err
}