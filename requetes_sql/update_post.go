package requetes_sql

import (
	uuid "github.com/satori/go.uuid"
	"log"
)

func UpdatePost(id uuid.UUID, title string, content string) {
	_, err := DB.Exec("UPDATE Posts SET Titre = ?, Contenu = ? WHERE id = ?", title, content, id)
	if err != nil {
		log.Fatal("UpdateUser: ", err)
	}
}
