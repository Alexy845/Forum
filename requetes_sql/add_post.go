package requetes_sql

import (
	uuid2 "forum/uuid"
	uuid "github.com/satori/go.uuid"
	"time"

	"log"
)

func AddPost(idAuteur uuid.UUID, titre string, contenu string, category uuid.UUID) {
	_, err := DB.Exec("INSERT INTO Posts (Id, Auteur, Titre, Contenu, Date, Likes, Dislikes, Category) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", uuid2.CreateUUID(), idAuteur, titre, contenu, time.Now().Format("2006-01-02 15:04:05"), 0, 0, category)
	if err != nil {
		log.Fatal(err)
	}
}
