package requetes_sql

import (
	"forum/structs"
	uuid2 "forum/uuid"
	"log"
	"time"
)

func AddComment(user structs.User, post structs.Post, contenu string) {
	_, err := DB.Exec("INSERT INTO Comments (Id, Auteur, Post, Contenu, Date) VALUES (?, ?, ?, ?, ?)", uuid2.CreateUUID(), user.Id, post.Id, contenu, time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		log.Fatal(err)
	}
}
