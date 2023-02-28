package requetes_sql

import (
	"forum/fonctions"
	uuid "github.com/satori/go.uuid"
)

func Verifuser(username string, password string) uuid.UUID {
	for _, user := range Getall() {
		if user.Username == username && fonctions.CompareHash(user.Password, password) {
			return user.Id
		}
	}
	return uuid.Nil
}
