package requetes_sql

import (
	"database/sql"
	"fmt"
	"forum/structs"
	uuid "github.com/satori/go.uuid"
	"log"
)

func GetAllLike() []structs.Like {
	req, err := DB.Query("SELECT Id, User, Post FROM Like")
	if err != nil {
		fmt.Println("Getall1")
		log.Fatal(err)
	}
	likes := []structs.Like{}
	for req.Next() {
		like := structs.Like{}
		user := uuid.UUID{}
		post := uuid.UUID{}
		err = req.Scan(&like.Id, &user, &post)
		if err != nil {
			fmt.Println("Getall2")
			log.Fatal(err)
		}
		like.User = GetUser(user)
		like.Post = GetPost(post)
		likes = append(likes, like)
	}
	defer func(req *sql.Rows) {
		err := req.Close()
		if err != nil {
			fmt.Println("Getall3")
			log.Fatal(err)
		}
	}(req)
	return likes
}
