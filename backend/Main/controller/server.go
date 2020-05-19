package controller

import (
	"fmt"
	"log"
	"time"
	"net/http"
	"encoding/json"
	"io/ioutil"

	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"github.com/gorilla/context"
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
	Username string 
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

	r.HandleFunc("/user", addUser).Methods("PUT")
	r.HandleFunc("/professions", getProfessions).Methods("GET")
	r.HandleFunc("/projectCategories", getProjectCategories).Methods("GET")
	r.HandleFunc("/login", auth).Methods("POST")

	secure.HandleFunc("/user", getUser).Methods("GET")
	secure.HandleFunc("/project", createProject).Methods("PUT")
	secure.HandleFunc("/project", getProjects).Methods("GET")
	secure.HandleFunc("/invitationProjectStudentTeam", genarateInvitationLink).Methods("GET")

	return http.ListenAndServe(Addr, handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(r))
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
			http.Error(w, "User not Found", http.StatusUnprocessableEntity)
			return
		}
		if err != nil {
			log.Println("Error finding the password for the username, ", credentials.Username)
			http.Error(w, "Error finding the password for the username", http.StatusInternalServerError)
			return
		}
		result := pkg.CompareHash(credentials.Password, hashPassword)
		if result != nil {
			log.Println(result)
			http.Error(w, "Wrong Password", http.StatusUnauthorized)
			return
		}
		exp := time.Now().Add(1 *time.Hour)
		claims := &claims{
			Username: credentials.Username,
			StandardClaims: jwt.StandardClaims{
				// In JWT, the expiry time is expressed as unix milliseconds
				ExpiresAt: exp.Unix(),
			},
		}
	
		
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
		tokenString, err := token.SignedString(jwtKey)

		if err != nil {
			log.Println("Error creating the token, " + err.Error())
			http.Error(w, "Error creating the token, " + err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write([]byte(tokenString))

}

//to authenticate the users
func isAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Token")
		if tokenString == ""{
			log.Println("No token Found")
			http.Error(w, "No token Provided", http.StatusUnauthorized)
			return
		}
		claims := &claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil {
			log.Println(err)
			if err == jwt.ErrSignatureInvalid {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}
			v, _ := err.(*jwt.ValidationError)
			if v.Errors == jwt.ValidationErrorExpired {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if !token.Valid {
			http.Error(w, "Token is not valid", http.StatusUnauthorized)
			return
		}
		context.Set(r, "username", claims.Username)
		next.ServeHTTP(w, r)
	})
}
 