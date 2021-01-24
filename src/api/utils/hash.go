package utils

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func CreateHash(passwd string) string {
	// Generate "hash" to store from user password
	hash, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
	if err != nil {
		// TODO: Properly handle error
		log.Fatal(err)
	}
	return string(hash)
}

func CompareHash(passwd, hash string) bool {
	// Comparing the password with the hash
	if err := bcrypt.CompareHashAndPassword([]byte(passwd), []byte(hash)); err != nil {
		return false
	}
	return true
}