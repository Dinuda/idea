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

	//add a new Investor
	r.HandleFunc("/addInvestor", func(w http.ResponseWriter, r *http.Request){
		var User models.Investor
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("ERR: Reading /addInvestor body" + err.Error())
		}

		err = json.Unmarshal(body, &User)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println("ERR: Unmarshal/addInvestor" + err.Error())
		}

	})

	return http.ListenAndServe(":8000", r)
}