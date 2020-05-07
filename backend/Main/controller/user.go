package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"

	"github.com/gorilla/mux"

	"../models"
)

//StartServer starts the server and listens
func StartServer()error{
	r := mux.NewRouter()

	//HomePage
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "hello World ")
	})

	//add a new entrepreneur
	r.HandleFunc("/addEntrepreneur", func(w http.ResponseWriter, r *http.Request){
		var User models.Entrepreneur
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("ERR: Reading /addEntrepreneur body" + err.Error())
		}

		err = json.Unmarshal(body, &User)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println("ERR: Unmarshal/addEntrepreneur" + err.Error())
		}

	})

	return http.ListenAndServe(":8000", r)
}