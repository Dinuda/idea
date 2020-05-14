package repository

import (
	"../models"
)

//GetProjectCategories get all the category from the db
func GetProjectCategories() ([]models.ProjectCategory, error) {
	result, err := selectProjectCatagoriesStmt.Query()
	if err != nil {
		return []models.ProjectCategory{}, err
	}
	var projectCategories []models.ProjectCategory
	for result.Next() {
		var projectCategory models.ProjectCategory
		err = result.Scan(&projectCategory.ID, &projectCategory.Name)
		if err != nil {
			return []models.ProjectCategory{}, err
		}
		projectCategories = append(projectCategories, projectCategory)
	}

	return projectCategories, err
}
