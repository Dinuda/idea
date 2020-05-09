package main

import(
	"log"


	"github.com/magiconair/properties"

	"./controller"
	"./repository"
)

var dbConfig repository.Config 
var serverConfig  controller.Config


func init(){
	log.Println("Initializing")
	p := properties.MustLoadFile("./config/config.properties", properties.UTF8)

	dbConfig = repository.Config{
		Username: p.MustGet("DB_USERNAME"),
		Password: p.MustGet("DB_PASS"),
		Protocol: p.MustGet("DB_PROTOCOL"),
		URL: p.MustGet("DB_URL"),
		Port: p.MustGet("DB_PORT"),
		Schema: p.MustGet("DB_SCHEMA"),
	}

	serverConfig = controller.Config{
		Host: p.MustGet("SERVER_HOST"),
		Port: p.MustGet("SERVER_PORT"),
	}
	//repository.Commit = p.MustGet("DB_COMMIT")
	
}

func main(){
	log.Println("Starting")
	var err error

	repository.DB, err = repository.Connect(dbConfig)
	if err != nil {
		log.Fatal("Error connecting to the server")
	}

	err = repository.Prepare()
	if err != nil{
		log.Fatal("Error preparing the Stmt ", err.Error())
	}

	err = controller.StartServer(serverConfig)
	if err != nil {
		log.Fatal("Error starting the Server", err.Error())
	}
}