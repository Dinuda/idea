package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"../models"
	"../pkg"
	"../repository"
	"../service"
)

//Addr used to configure the server
var Addr string

//StartServer starts the server and listens
func StartServer() error {
	r := mux.NewRouter().StrictSlash(true)
	secure := r.PathPrefix("/auth").Subrouter()
	secure.Use(isAuth)
	//HomePage
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello World ")
	})

	//add a new User
	r.HandleFunc("/addUser", func(w http.ResponseWriter, r *http.Request) {
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

	})

	//Get all the professions
	r.HandleFunc("/getProfessions", func(w http.ResponseWriter, r *http.Request) {
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

	}).Methods("GET")

	//Get all the Categories
	r.HandleFunc("/getCategories", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Getting Categories")
		var projectCategories []models.ProjectCategory
		var err error
		projectCategories, err = service.GetProjectCategories()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if err := json.NewEncoder(w).Encode(projectCategories); err != nil {
			log.Println("Error responding to the /getProjectCategories")
		}

	}).Methods("GET")

	secure.HandleFunc("/getUser", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "authenticated")
	})

	return http.ListenAndServe(Addr, r)
}

//to authenticate the users
func isAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username := r.Header.Get("username")
		password := r.Header.Get("password")
		if username == "" || password == "" {
			fmt.Fprintf(w, "User or Password not given")
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}
		fmt.Println(username, password)
		hashPassword, err := repository.GetUserPassword(username)
		if hashPassword == "" {
			w.WriteHeader(http.StatusUnprocessableEntity)
			fmt.Fprintf(w, "User not Found")
			return
		}
		if err != nil {
			log.Println("Error finding the password for the username, ", username)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		result := pkg.CompareHash(password, hashPassword)
		if result != nil {
			log.Println(result)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
