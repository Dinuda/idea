package controller

import (
	"github.com/gorilla/mux"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"../models"
	"../service"
)

//gets all the categories of the projects
func getProjectCategories(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting Categories")
	projectCategories, err := service.GetProjectCategories()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(projectCategories); err != nil {
		log.Println("Error responding to the /getProjectCategories")
	}

}

func createProject(w http.ResponseWriter, r *http.Request) {
	log.Println("Creating a new Project")
	var project models.Project
	username := r.Header.Get("username")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading the body of /createProject, " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Fprintf(w, "Error reading the body of /createProject, "+err.Error())
		return
	}

	err = json.Unmarshal(body, &project)
	if err != nil {
		log.Println("Error unmarshelling /createProject, " + err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Fprintf(w, "Error unmarshelling /createProject, "+err.Error())
		return
	}

	project, err = service.CreateProject(project, username)
	if err != nil {
		log.Println("Error creating a new project /createProject, " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Fprintf(w, "Error creating a new project /createProject,"+err.Error())
		return
	}
	json.NewEncoder(w).Encode(project)
}


func addStudentToTeam(w http.ResponseWriter, r *http.Request){
	log.Println("Adding a student to a team")
	username := r.Header.Get("username")
	teamID := mux.Vars(r)["teamID"]
	intTeamID, _ := strconv.Atoi(teamID)
	err := service.AddStudentToTeam(username, intTeamID)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Successfully added the student to the team")
}

func getProjects(w http.ResponseWriter, r *http.Request){
	log.Println("Getting Projects")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading the request of /projects, " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var user models.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Println("Error reading the request of /projects, " + err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	projects, err := service.GetProjects(user.Username)
	if err != nil {
		log.Println("Error getting the Projects, " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(projects)
	if err != nil {
		log.Println("Error sendin the response to /projects, " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}