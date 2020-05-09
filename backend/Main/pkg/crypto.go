package pkg

import (
	"log"
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"
)



//GenarateHash is used to genarate hash keys from user passwords to store in the DB
func GenarateHash(pass string)(string, error){
	log.Println("Hasing a password")
	bytePass := []byte(pass)
	hashPass, err := bcrypt.GenerateFromPassword(bytePass, bcrypt.DefaultCost)
	return hex.EncodeToString(hashPass), err
}

//CompareHash is used to compare a string and a Hash
func CompareHash(pass string, hash string)(error){
	log.Println("Comparing a password")
	bytePass := []byte(pass)
	byteHash := []byte(hash)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePass)
	return err
}

