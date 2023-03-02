package requetes_sql

import (
	"database/sql"
	"forum/structs"
	uuid "github.com/satori/go.uuid"
	"log"
)

func GetPost(id uuid.UUID) structs.Post {
	req, err := DB.Query("SELECT Id, Auteur, Contenu, Titre, Date, Likes FROM Posts where Id = ?", id)
	if err != nil {
		log.Fatal()
	}
	post := structs.Post{}
	for req.Next() {
		err = req.Scan(&post.Id, &post.Auteur, &post.Contenu, &post.Titre, &post.Date, &post.Likes)
		if err != nil {
			log.Fatal(err)
		}
	}
	defer func(req *sql.Rows) {
		err := req.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(req)
	return post
}
