package requetes_sql

import (
	"database/sql"
	"forum/structs"
	uuid "github.com/satori/go.uuid"
	"log"
)

func GetAllPosts() []structs.Post {
	req, err := DB.Query("SELECT Id, Auteur, Contenu, Titre, Date, Likes, Dislikes, Category FROM Posts")
	if err != nil {
		log.Fatal(err)
	}
	posts := []structs.Post{}
	for req.Next() {
		post := structs.Post{}
		auteur := uuid.UUID{}
		category := uuid.UUID{}
		err = req.Scan(&post.Id, &auteur, &post.Contenu, &post.Titre, &post.Date, &post.Likes, &post.Dislikes, &category)
		if err != nil {
			log.Fatal(err)
		}
		if err != nil {
			log.Fatal(err)
		}
		post.Category = GetCategory(category)
		post.Auteur = GetUser(auteur)
		posts = append(posts, post)
	}
	defer func(req *sql.Rows) {
		err := req.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(req)
	return posts
}
