package controller

import (
	"log"
	"encoding/json"
	"net/http"
	

	"../service"
)

//gets all the categories of the projects
func getProjectCategories(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting Categories")
	projectCategories, err := service.GetProjectCategories()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(projectCategories); err != nil {
		log.Println("Error responding to the /getProjectCategories")
	}

}