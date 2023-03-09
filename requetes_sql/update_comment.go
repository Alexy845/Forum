package requetes_sql

import (
	uuid "github.com/satori/go.uuid"
	"log"
)

func UpdateComment(id uuid.UUID, content string) {
	_, err := DB.Exec("UPDATE Comments SET Contenu = ? WHERE id = ?", content, id)
	if err != nil {
		log.Fatal("UpdateUser: ", err)
	}
}
