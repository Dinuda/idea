package service

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)
//StartServer starts the rest api
func StartServer(addr string, port int)error{
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Hello world")
	})
	return nil 
}