package repository

import (
	"fmt"

	"../models"
)

//GetProjectCategories get all the category from the db
func GetProjectCategories() ([]models.ProjectCategory, error) {
	result, err := selectProjectCatagoriesStmt.Query()
	if err != nil {

		return []models.ProjectCategory{}, fmt.Errorf("Error getting ProjectCategories, "+ err.Error())
	}
	var projectCategories []models.ProjectCategory
	for result.Next() {
		var projectCategory models.ProjectCategory
		err = result.Scan(&projectCategory.ID, &projectCategory.Name)
		if err != nil {
			return []models.ProjectCategory{}, fmt.Errorf("Error getting ProjectCategories, "+ err.Error())
		}
		projectCategories = append(projectCategories, projectCategory)
	}

	return projectCategories, nil
}


//CreateStudentTeam creates a new student team
func CreateStudentTeam(userID int)(int,error){
	result, err := createStudentTeamStmt.Exec(userID)
	if err != nil {
		return 0, fmt.Errorf("Error inserting a student to the student team, " + err.Error())
	}
	var lastInsertID int64
	lastInsertID, err = result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("Error getting the last insert ID, "+ err.Error())
	}
	return int(lastInsertID), nil
}


//CreateProject creates a new Project
func CreateProject(project models.Project) (int, error){
	result, err := insertProjectStmt.Exec(
		project.Title,
		project.Description,
		project.CreatedDate,
		project.Category,
		project.StudentTeamID,
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
