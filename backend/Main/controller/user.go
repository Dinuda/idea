package controller

import (
	"log"
	"net/http"
	"encoding/json"
	"io/ioutil"

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
		log.Println("ERR: Unmarshal /addUser, " + err.Error())
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
	setupResponse(&w, r)
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
	var user models.User
	body, err := ioutil.ReadAll(r.Body)
	//fmt.Println(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("ERR: Reading /getUser body, " + err.Error())
		return
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("ERR: Unmarshal /getUser, " + err.Error())
		return
	}
	if user.Username == ""{
		w.WriteHeader(http.StatusBadRequest)
		//fmt.Fprintf(w, "No username provided")
	}
	log.Println(user.Username)
	user, err = service.GetUser(user.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	if err := json.NewEncoder(w).Encode(user); err != nil {
		log.Println("Error responding to /GetUser ", err.Error())
	}

}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	log.Println("setting CORS")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
    (*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}