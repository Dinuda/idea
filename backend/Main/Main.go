package main

import (
	"log"

	"github.com/magiconair/properties"

	"./controller"
	"./repository"
)

func init() {
	log.Println("Initializing")
	p := properties.MustLoadFile("./config/config.properties", properties.UTF8)

	repository.DNS = p.MustGet("DB_DNS")
	controller.Addr = p.MustGet("SERVER_ADDR")

}

func main() {
	log.Println("Starting server on", controller.Addr)
	var err error

	repository.DB, err = repository.Connect()
	if err != nil {
		log.Fatal("Error connecting to the server,"+err.Error())
	}

	err = repository.Prepare()
	if err != nil {
		log.Fatal("Error preparing the Stmt ", err.Error())
	}
	log.Println("Successfully prepared the Stmt")

	err = controller.StartServer()
	if err != nil {
		log.Fatal("Error starting the Server", err.Error())
	}
}
