package requetes_sql

import (
	"database/sql"
	"fmt"
	"forum/structs"
	"log"
)

func GetAllPosts() []structs.Post {
	req, err := DB.Query("SELECT Id, Auteur, Contenu, Titre, Date FROM Posts")
	if err != nil {
		fmt.Println("Getall1")
		log.Fatal(err)
	}
	posts := []structs.Post{}
	for req.Next() {
		post := structs.Post{}
		err = req.Scan(&post.Id, &post.Auteur, &post.Contenu, &post.Titre, &post.Date)
		if err != nil {
			fmt.Println("Getall2")
			log.Fatal(err)
		}
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
