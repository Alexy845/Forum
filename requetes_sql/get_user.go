package requetes_sql

import (
	"database/sql"
	"forum/structs"
	"log"
)

func GetUser(id int) structs.User {
	req, err := DB.Query("SELECT id, pseudo, prenom, nom, date_membre, mdp, email FROM User where Id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	defer func(req *sql.Rows) {
		err := req.Close()
		if err != nil {

		}
	}(req)
	req.Next()
	user := structs.User{}
	err = req.Scan(&user.Id, &user.Username, &user.FirstName, &user.LastName, &user.CreationDate, &user.Password, &user.Email)
	if err != nil {
		log.Fatal(err)
	}
	return user
}
