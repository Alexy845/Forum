package requetes_sql

import (
	"database/sql"
	"forum/structs"
	uuid "github.com/satori/go.uuid"
	"log"
)

func GetAllComment() []structs.Comment {
	req, err := DB.Query("SELECT * FROM Comments")
	if err != nil {
		log.Fatal(err)
	}
	comments := []structs.Comment{}
	for req.Next() {
		comment := structs.Comment{}
		user := uuid.UUID{}
		post := uuid.UUID{}
		err = req.Scan(&comment.Id, &user, &post, &comment.Contenu, &comment.Date)
		if err != nil {
			log.Fatal(err)
		}
		comment.Auteur = GetUser(user)
		comment.Post = GetPost(post)
		comments = append(comments, comment)
	}
	defer func(req *sql.Rows) {
		err := req.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(req)
	return comments
}
