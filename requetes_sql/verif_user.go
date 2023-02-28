package requetes_sql

import "forum/fonctions"

func Verifuser(username string, password string) int {
	for _, user := range Getall() {
		if user.Username == username && fonctions.CompareHash(user.Password, password) {
			return user.Id
		}
	}
	return 0
}
