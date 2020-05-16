package controller

import (
	"fmt"
	"log"
	"time"
	"net/http"
	"encoding/json"
	"io/ioutil"

	"github.com/gorilla/mux"
	jwt "github.com/dgrijalva/jwt-go"

	"../pkg"
	"../repository"
)
//Credentials to keep track of credentials
type credentials struct {
	Username string 
	Password string
}

//Claims of the token
type claims struct {
	username string 
	jwt.StandardClaims
}

var jwtKey = []byte(`MIICXAIBAAKBgHTSnN3g0neR4kPUiPRA3Qb6T7zbSh31uvRcKsSR5lF5mprnYr6a
	Q8VwSyj/gimUZk7z4zeNkl5yyJCmzUxeOAs/Xt+sq4yscqwik1EXwTyZGm0e45MW
	2/h4PEUFELMAUtqy20HpUKXuKzNZsz20bdTS1pgA+hN33Uib68cYQNmXAgMBAAEC
	gYBHl6bApv30fve9/+rqXTHXC+F/6Je0YppvFGi1TIBsX+yqj7DJBDsSLW4yMtuu
	5Z4JNpeBvQX9UbSuSTq5WWhYT5X28EDlimlwu0n/TUC7CrE1JMf6Je7HSB3IfgAw
	MMRShybWY+kX0etfx0j/1oVHilHrwkkJJU9cJn4wtpYHWQJBALuIBp8T99zFJmdv
	i1rQ08p97+oawTjx6hPs+m3u3DLv/VObPuigPcKLgBqpWDeEUe0d7aA+jC6RQpkD
	J3Ibw1sCQQCfeatEXx136FjsX9tJetNsNr94e3IpU2rhLs/T98nU2M7N7SIYx6K1
	meOVKNdvr2/rBYeHnYrT/HW/vK4GG4N1AkEAtGl7vUTPmwPMG4yTK25lopQf4D+X
	DjqlsD+2+VXnX9XEB8/96HxojiX4uy2Z4ecZjh3Rwu0Jna8/u8buBvgwqwJAL4kl
	2wB7GTXh47uC8vkwsi3zjudFFTpvPmYkvus6dz6VDl7j7fz77CPN6bU92mWx950z
	U+JK8ntrYdbNDLcAzQJBALdfXC7Ef2+mYaB7ioy6kUxcjo4ubRryQCdYRm9s7VOd`)

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

	r.HandleFunc("/addUser", addUser).Methods("PUT")
	r.HandleFunc("/getProfessions", getProfessions).Methods("GET")
	r.HandleFunc("/getProjectCategories", getProjectCategories).Methods("GET")
	r.HandleFunc("/login", auth).Methods("POST")

	secure.HandleFunc("/getUser", getUser).Methods("GET")
	secure.HandleFunc("/createProject", createProject).Methods("POST")

	return http.ListenAndServe(Addr, r)
}

func auth(w http.ResponseWriter, r *http.Request){
	var credentials credentials
	
	body, err := ioutil.ReadAll(r.Body)
	//fmt.Println(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("ERR: Reading /auth body, " + err.Error())
		return
	}

	err = json.Unmarshal(body, &credentials)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("ERR: Unmarshal /auth, " + err.Error())
		return
	}
	//fmt.Println(r.Body)
	hashPassword, err := repository.GetUserPassword(credentials.Username)
		if hashPassword == "" {
			w.WriteHeader(http.StatusUnprocessableEntity)
			fmt.Fprintf(w, "User not Found")
			return
		}
		if err != nil {
			log.Println("Error finding the password for the username, ", credentials.Username)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		result := pkg.CompareHash(credentials.Password, hashPassword)
		if result != nil {
			log.Println(result)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		exp := time.Now().Add(1 *time.Hour)

		claims := &claims{
			username: credentials.Username,
			StandardClaims: jwt.StandardClaims{
				// In JWT, the expiry time is expressed as unix milliseconds
				ExpiresAt: exp.Unix(),
			},
		}
	
		
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
		tokenString, err := token.SignedString(jwtKey)

		if err != nil {
			log.Println("Error creating the token, " + err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(w).Encode(tokenString)
		if err != nil {
			log.Println("Error sending the token, " + err.Error())
			w.WriteHeader(http.StatusInternalServerError)
		}

}

//to authenticate the users
func isAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Token")
		if tokenString == ""{
			log.Println("No token Found")
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "No token provided")
			return
		}
		claims := &claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil {
			log.Println(err)
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		fmt.Println("hi")
		next.ServeHTTP(w, r)
	})
}
 