package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"../models"
	"../service"
)

//Config used to configure the server
type Config struct {
	Host string
	Port string
}

//StartServer starts the server and listens
func StartServer(config Config) error {
	r := mux.NewRouter()

	//HomePage
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello World ")
	})

	//add a new User
	r.HandleFunc("/addUser", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Adding new User called")
		var user models.User
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("ERR: Reading /addUser body" + err.Error())
		}

		err = json.Unmarshal(body, &user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println("ERR: Unmarshal/addUser" + err.Error())
		}
		//add a new user
		service.AddUser(user)

	})
	path := config.Host + ":" + config.Port

	return http.ListenAndServe(path, r)
}
