package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main(){
	r := mux.NewRouter()
	r.HandleFunc("/" func(w http.ResponseWriter, r *http.Request){
		fmt.Println("Hello world")
		fmt.Fprintf(w, "Hello world")
	})
}