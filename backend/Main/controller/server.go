package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"../pkg"
	"../repository"
)

//Addr used to configure the server
var Addr string

//StartServer starts the server and listens
func StartServer() error {
	r := mux.NewRouter().StrictSlash(true)
	secure := r.PathPrefix("/auth").Subrouter()
	secure.Use(isAuth)
	
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello World ")
	})
	
	r.HandleFunc("/addUser", addUser).Methods("POST")
	r.HandleFunc("/getProfessions", getProfessions).Methods("GET")
	r.HandleFunc("/getProjectCategories", getProjectCategories).Methods("GET")

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
