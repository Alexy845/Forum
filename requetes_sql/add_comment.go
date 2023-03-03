package requetes_sql

import (
	uuid2 "forum/uuid"
	uuid "github.com/satori/go.uuid"
	"log"
	"time"
)

func AddComment(user uuid.UUID, post uuid.UUID, contenu string) {
	_, err := DB.Exec("INSERT INTO Comments (Id, Auteur, Post, Contenu, Date) VALUES (?, ?, ?, ?, ?)", uuid2.CreateUUID(), user, post, contenu, time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		log.Fatal(err)
	}
}
