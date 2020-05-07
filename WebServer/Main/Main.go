package main

import (
	"fmt"
	"html/template"
	"net/http"
	"log"

	"github.com/gorilla/mux"
)
//templates
var(
	homePage *template.Template
)

type pageInfo struct{
	Title string
}

func loadPages(){
	homePage = template.Must(template.ParseFiles("./templates/main.html"))
}

func router()error{
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		homePage.Execute(w, pageInfo{Title: "this is a test"})
		fmt.Println("Hello world")
		//fmt.Fprintf(w, "Hello world")
	})

	//Start the server 
	return http.ListenAndServe(":8080", r)
}

func main() {
	log.Fatal(router())
}
