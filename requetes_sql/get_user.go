package requetes_sql

import (
	"database/sql"
	"fmt"
	"forum/structs"
	uuid "github.com/satori/go.uuid"
	"log"
)

func GetUser(id uuid.UUID) structs.User {
	req, err := DB.Query("SELECT id, pseudo, prenom, nom, date_membre, mdp, email FROM User where Id = ?", id)
	if err != nil {
		fmt.Println("GetUser1")
		log.Fatal(err)
	}
	user := structs.User{}
	for req.Next() {
		err = req.Scan(&user.Id, &user.Username, &user.FirstName, &user.LastName, &user.CreationDate, &user.Password, &user.Email)
		if err != nil {
			fmt.Println("GetUser2")
			log.Fatal(err)
		}
	}
	defer func(req *sql.Rows) {
		err := req.Close()
		if err != nil {
			fmt.Println("GetUser3")
			log.Fatal(err)
		}
	}(req)
	return user
}
