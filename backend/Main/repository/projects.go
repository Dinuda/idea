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

//CreateStudentTeam creates a new student team
func CreateStudentTeam(userID int) (int, error) {
	result, err := createStudentTeamStmt.Exec(userID)
	if err != nil {
		return 0, fmt.Errorf("Error inserting a student to the student team, " + err.Error())
	}
	var lastInsertID int64
	lastInsertID, err = result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("Error getting the last insert ID, " + err.Error())
	}
	return int(lastInsertID), nil
}

//CreateProject creates a new Project
func CreateProject(project models.Project) (int, error) {
	result, err := insertProjectStmt.Exec(
		project.Title,
		project.Description,
		project.Category,
		project.StudentTeamID,
		project.HostID,
	)
	if err != nil {
		return 0, fmt.Errorf("Error creating a project, " + err.Error())
	}
	var rowsAffected int64
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("Error getting no of rows Affected," + err.Error())
	}
	return int(rowsAffected), nil
}

//AddStudentToTeam adds a new student to the studentteam of the project
func AddStudentToTeam(teamID int, userID int) (int, error) {
	result, err := insertStudentToTeamStmt.Exec(teamID, userID)
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
func GetProjects(hostID int)([]models.Project, error){
	var projects []models.Project
	result, err := selectProjectStmt.Query(hostID)
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
			&project.ClosedDate,
			&project.StudentTeamID,
			&project.InvestorTeamID,
			&project.Category,
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