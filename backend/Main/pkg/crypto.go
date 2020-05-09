package pkg

import (
	"log"
	"golang.org/x/crypto/bcrypt"
)



//GenarateHash is used to genarate hash keys from user passwords to store in the DB
func GenarateHash(pass string)([]byte, error){
	log.Println("Hasing a password")
	bytePass := []byte(pass)
	return bcrypt.GenerateFromPassword(bytePass, bcrypt.DefaultCost)
}

