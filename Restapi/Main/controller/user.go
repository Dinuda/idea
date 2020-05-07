package controller

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//StartServer starts the server and listens
func StartServer()error{
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "hello World ")
	})
}