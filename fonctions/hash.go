package fonctions

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func Hash(str string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Hash: ", err)
	}
	return string(hash)
}

func CompareHash(hash string, str string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(str))
	if err != nil {
		return false
	}
	return true
}
