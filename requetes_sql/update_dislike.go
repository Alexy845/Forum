package requetes_sql

import (
	uuid "github.com/satori/go.uuid"
	"log"
)

func UpdateDislike(id uuid.UUID, dislike int) {
	_, err := DB.Exec("UPDATE Posts SET Dislikes = ? WHERE id = ?", dislike, id)
	if err != nil {
		log.Fatal("UpdateMDP: ", err)
	}
}
