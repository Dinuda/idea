package pkg

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

//GenarateHash is used to genarate hash keys from user passwords to store in the DB
func GenarateHash(password string) (string, error) {
	log.Println("Hasing a password")
	bytePassword := []byte(password)
	hashPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	log.Println(string(hashPassword))
	return string(hashPassword), err
}

//CompareHash is used to compare a string and a Hash
func CompareHash(password string, hash string) error {
	log.Println("Comparing a password")
	bytePassword := []byte(password)
	byteHash := []byte(hash)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePassword)
	return err
}
