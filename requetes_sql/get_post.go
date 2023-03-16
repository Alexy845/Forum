package requetes_sql

import (
	"database/sql"
	"forum/structs"
	uuid "github.com/satori/go.uuid"
	"log"
)

func GetPost(id uuid.UUID) structs.Post {
	req, err := DB.Query("SELECT Id, Auteur, Contenu, Titre, Date, Likes, Dislikes FROM Posts where Id = ?", id)
	if err != nil {
		log.Fatal()
	}
	post := structs.Post{}
	user := uuid.UUID{}
	for req.Next() {
		err = req.Scan(&post.Id, &user, &post.Contenu, &post.Titre, &post.Date, &post.Likes, &post.Dislikes)
		if err != nil {
			log.Fatal(err)
		}
		post.Auteur = GetUser(user)
	}
	defer func(req *sql.Rows) {
		err := req.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(req)
	return post
}
