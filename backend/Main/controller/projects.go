package controller

import (
	"github.com/gorilla/mux"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"../models"
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

func createProject(w http.ResponseWriter, r *http.Request) {
	log.Println("Creating a new Project")
	var project models.Project
	username := mux.Vars(r)["username"]
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading the body of /createProject, " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error reading the body of /createProject, "+err.Error())
		return
	}

	err = json.Unmarshal(body, &project)
	if err != nil {
		log.Println("Error unmarshelling /createProject, " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error unmarshelling /createProject, "+err.Error())
		return
	}

	project, err = service.CreateProject(project, username)
	if err != nil {
		log.Println("Error creating a new project /createProject, " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error creating a new project /createProject,"+err.Error())
		return
	}
	json.NewEncoder(w).Encode(project)
}
