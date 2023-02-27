package requetes_sql

import (
	"forum/structs"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func Getall() []structs.User {
	req, err := DB.Query("SELECT id, pseudo, prenom, nom, date_membre, mdp, email FROM User")
	if err != nil {
		log.Fatal(err)
	}
	defer req.Close()
	users := []structs.User{}
	for req.Next() {
		user := structs.User{}
		err = req.Scan(&user.Id, &user.Username, &user.FirstName, &user.LastName, &user.CreationDate, &user.Password, &user.Email)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	return users
}
