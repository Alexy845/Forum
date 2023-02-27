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
