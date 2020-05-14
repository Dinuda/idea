package controller

import (
	"log"
	"net/http"
	"encoding/json"
	"io/ioutil"

	"github.com/gorilla/mux"

	"../models"
	"../service"
)

//adds a new user
func addUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Adding new User called")
	var user models.User
	body, err := ioutil.ReadAll(r.Body)
	//fmt.Println(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("ERR: Reading /addUser body, " + err.Error())
		return
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("ERR: Unmarshal/addUser, " + err.Error())
		return
	}
	log.Println("hi", user)
	//add a new user
	err = service.AddUser(user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusAlreadyReported)
		return
	}
	log.Println("Adding User Successfull")

}

//gets all the professions
func getProfessions(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting Professions")
	//var professions []models.Profession
	//var err error
	professions, err := service.GetProfessions()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(professions); err != nil {
		log.Println("Error responding to the /getProfessions")
	}

}

func getUser(w http.ResponseWriter, r *http.Request){
	log.Println("Getting User Info")
	username := mux.Vars(r)["username"]
	user, err := service.GetUser(username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	if err := json.NewEncoder(w).Encode(user); err != nil {
		log.Println("Error responding to /GetUser ", err.Error())
	}

}