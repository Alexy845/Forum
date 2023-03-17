package requetes_sql

import (
	"database/sql"
	"forum/structs"
	uuid "github.com/satori/go.uuid"
	"log"
)

func SearchPostByCategory(category string) []structs.Post {
	req, err := DB.Query("SELECT Posts.Id from Posts join Category C on Posts.Category = C.Id where C.Name LIKE ?", category)
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
