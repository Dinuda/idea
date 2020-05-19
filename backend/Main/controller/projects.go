package controller

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/context"

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
	username := context.Get(r, "username").(string)
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
	err = json.NewEncoder(w).Encode(project)
	if err != nil {
		log.Println("Error sending the token, " + err.Error())
		http.Error(w, "Error sending the token, " + err.Error(), http.StatusInternalServerError)
	}
}


func addStudentToTeam(w http.ResponseWriter, r *http.Request){
	log.Println("Adding a student to a team")
	username := context.Get(r, "username").(string)
	projectID := mux.Vars(r)["projectID"]
	intProjectID, _ := strconv.Atoi(projectID)
	err := service.AddStudentToProjectStudentTeam(username, intProjectID)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Successfully added the student to the team")
}

func getProjects(w http.ResponseWriter, r *http.Request){
	log.Println("Getting Projects")
	username := context.Get(r, "username").(string)

	projects, err := service.GetProjects(username)
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

func genarateInvitationLink(w http.ResponseWriter, r *http.Request){
	log.Println("Genarating Invitation link")
	projectID := r.Header.Get("projectID")
	if projectID == ""{
		log.Println("Error no username or the projectID given")
		http.Error(w, "Error no username or the projectID given", http.StatusBadRequest)
		return
	}
	intProjectID, _ := strconv.Atoi(projectID)
	token, err := service.GenarateProjectInvitationCode(intProjectID)
	if err != nil {
		log.Println("Error Genarating the code" + err.Error())
		http.Error(w, "Error Genarating the code" + err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(token))
}