package requetes_sql

import (
	"database/sql"
	"forum/structs"
	uuid "github.com/satori/go.uuid"
	"log"
)

func GetPostByUser(UserId uuid.UUID) []structs.Post {
	req, err := DB.Query("SELECT Id FROM Posts where Posts.Auteur = ?", UserId)
	if err != nil {
		log.Fatal(err)
	}
	var posts []structs.Post
	for req.Next() {
		id := uuid.UUID{}
		err = req.Scan(&id)
		if err != nil {
			log.Fatal(err)
		}
		if err != nil {
			log.Fatal(err)
		}
		posts = append(posts, GetPost(id))
	}
	defer func(req *sql.Rows) {
		err := req.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(req)
	return posts
}
