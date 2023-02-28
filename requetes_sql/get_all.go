package requetes_sql

import (
	"database/sql"
	"fmt"
	"forum/structs"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func Getall() []structs.User {
	req, err := DB.Query("SELECT id, pseudo, prenom, nom, date_membre, mdp, email FROM User")
	if err != nil {
		fmt.Println("Getall1")
		log.Fatal(err)
	}
	users := []structs.User{}
	for req.Next() {
		user := structs.User{}
		err = req.Scan(&user.Id, &user.Username, &user.FirstName, &user.LastName, &user.CreationDate, &user.Password, &user.Email)
		if err != nil {
			fmt.Println("Getall2")
			log.Fatal(err)
		}
		users = append(users, user)
	}
	defer func(req *sql.Rows) {
		err := req.Close()
		if err != nil {
			fmt.Println("Getall3")
			log.Fatal(err)
		}
	}(req)
	return users
}
