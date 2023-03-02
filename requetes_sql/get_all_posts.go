package requetes_sql

import (
	"database/sql"
	"fmt"
	"forum/structs"
	uuid "github.com/satori/go.uuid"
	"log"
)

func GetAllPosts() []structs.Post {
	req, err := DB.Query("SELECT Id, Auteur, Contenu, Titre, Date, Likes FROM Posts")
	if err != nil {
		fmt.Println("Getall1")
		log.Fatal(err)
	}
	posts := []structs.Post{}
	for req.Next() {
		post := structs.Post{}
		auteur := uuid.UUID{}
		err = req.Scan(&post.Id, &auteur, &post.Contenu, &post.Titre, &post.Date, &post.Likes)
		if err != nil {
			fmt.Println("Getall2")
			log.Fatal(err)
		}
		post.Auteur = GetUser(auteur)
		posts = append(posts, post)
	}
	defer func(req *sql.Rows) {
		err := req.Close()
		if err != nil {
			fmt.Println("Getall3")
			log.Fatal(err)
		}
	}(req)
	return posts
}
