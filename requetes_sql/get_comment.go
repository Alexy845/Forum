package requetes_sql

import (
	"database/sql"
	"forum/structs"
	uuid "github.com/satori/go.uuid"
	"log"
)

func GetComment(id uuid.UUID) structs.Comment {
	req, err := DB.Query("SELECT * FROM Comments where Id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	comment := structs.Comment{}
	for req.Next() {
		user := uuid.UUID{}
		post := uuid.UUID{}
		err = req.Scan(&comment.Id, &user, &post, &comment.Contenu, &comment.Date)
		if err != nil {
			log.Fatal(err)
		}
		comment.Auteur = GetUser(user)
		comment.Post = GetPost(post)
	}
	defer func(req *sql.Rows) {
		err := req.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(req)
	return comment
}
