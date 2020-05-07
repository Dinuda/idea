package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)
//templates
type Pages struct{
	homePage *template.Template
}

func loadPages(pages Pages)*Pages{
	Pages.homePage = template.Must(template.ParseFiles("../templates/main.vue"))
}

func router(pages Pages){
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		pages.homePage.Execute(w, nil)
		fmt.Println("Hello world")
		//fmt.Fprintf(w, "Hello world")
	})
}

func main() {
	loadPages(Pages{nil})
	
}
