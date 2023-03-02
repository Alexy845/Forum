package requetes_sql

import (
	"forum/structs"
	uuid2 "forum/uuid"
	"log"
)

func AddLike(post structs.Post, user structs.User) {
	_, err := DB.Exec("INSERT INTO Like (Id, User, Post) VALUES (?, ?, ?)", uuid2.CreateUUID(), user.Id, post.Id)
	if err != nil {
		log.Fatal(err)
	}
}
