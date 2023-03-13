package requetes_sql

import (
	"database/sql"
	"forum/structs"
	uuid "github.com/satori/go.uuid"
	"log"
)

func SearchUser(input string) []structs.User {
	req, err := DB.Query("SELECT Id FROM User where User.Pseudo LIKE ?", input)
	if err != nil {
		log.Fatal(err)
	}
	users := []structs.User{}
	for req.Next() {
		id := uuid.UUID{}
		err = req.Scan(&id)
		if err != nil {
			log.Fatal(err)
		}
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, GetUser(id))
	}
	defer func(req *sql.Rows) {
		err := req.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(req)
	return users
}
