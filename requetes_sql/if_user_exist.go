package requetes_sql

import (
	"fmt"
)

func IfUserExist(email string, username string, firstname string, lastname string, password string) bool {
	for i := 0; i < GetTailleUser(); i++ {
		if Getemail(i) == email {
			fmt.Println("Cette adresse email est déjà utilisée")
			return true
		} else if GetUser(i).Username == username {
			fmt.Println("Ce pseudo est déjà utilisé")
			return true
		}
	}
	return false
}
