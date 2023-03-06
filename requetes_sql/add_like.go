package requetes_sql

import (
	uuid2 "forum/uuid"
	uuid "github.com/satori/go.uuid"
	"log"
)

func AddLike(post uuid.UUID, user uuid.UUID) {
	_, err := DB.Exec("INSERT INTO Like (Id, User, Post) VALUES (?, ?, ?)", uuid2.CreateUUID(), user, post)
	if err != nil {
		log.Fatal(err)
	}
}
