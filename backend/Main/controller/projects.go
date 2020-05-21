package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
	
	//"github.com/gorilla/mux"
	"github.com/gorilla/context"
	jwt "github.com/dgrijalva/jwt-go"

	"../models"
	"../service"
)

//gets all the categories of the projects
func getProjectCategories(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting Categories")
	projectCategories, err := service.GetProjectCategories()
	if err != nil || projectCategories == nil{
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
	log.Println("Adding a Student to a ProjectStudentTeam")
	var objmap map[string]json.RawMessage
	username := context.Get(r, "username").(string)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading the body /addStudentToTeam " + err.Error())
		http.Error(w, "Error reading the body /addStudentToTeam" + err.Error(), http.StatusInternalServerError)
	}
	err = json.Unmarshal(body, &objmap)
	if err != nil {
		log.Println("Error unmarshalling the body /addStudentToTeam " + err.Error())
		http.Error(w, "Error unmarshalling the body /addStudentToTeam" + err.Error(), http.StatusBadRequest)
	}
	var tokenString string
	err = json.Unmarshal(objmap["token"], &tokenString)
	if err != nil {
		log.Println("Error unmarshalling the body /addStudentToTeam " + err.Error())
		http.Error(w, "Error unmarshalling the body /addStudentToTeam" + err.Error(), http.StatusBadRequest)
	}
	if tokenString == ""{
		log.Println("No token Found")
		http.Error(w, "No token Provided", http.StatusUnauthorized)
		return
	}
	fmt.Println(tokenString)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		log.Println(err)
		if err == jwt.ErrSignatureInvalid {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		v, _ := err.(*jwt.ValidationError)
		if v.Errors == jwt.ValidationErrorExpired {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims)
		err := service.AddStudentToProjectStudentTeam(username, claims["ProjectID"].(int))
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "Successfully added the student to the team")
		return
	}

	http.Error(w, "Token is not valid or claims not available", http.StatusUnauthorized)
	
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
		log.Println("Error projectID given not given")
		http.Error(w, "Error projectID not given", http.StatusBadRequest)
		return
	}
	intProjectID, _ := strconv.Atoi(projectID)
	exp := time.Now().Add(48 *time.Hour)
	
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["projectID"] = intProjectID
		claims["exp"] = exp.Unix()
	
		tokenString, err := token.SignedString(jwtKey)

		if err != nil {
			log.Println("Error creating the token, " + err.Error())
			http.Error(w, "Error creating the token, " + err.Error(), http.StatusInternalServerError)
			return
		}
		
	
	w.Write([]byte(tokenString))
}