package requetes_sql

import (
	"forum/hash"
	uuid "github.com/satori/go.uuid"
)

func Verifuser(username string, password string) uuid.UUID {
	for _, user := range GetallUsers() {
		if user.Username == username && hash.CompareHash(user.Password, password) {
			return user.Id
		}
	}
	return uuid.Nil
}
