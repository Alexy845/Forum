package requetes_sql

import (
	uuid "github.com/satori/go.uuid"
	"log"
)

func UpdateLike(id uuid.UUID, like int) {
	_, err := DB.Exec("UPDATE Posts SET Likes = ? WHERE id = ?", like, id)
	if err != nil {
		log.Fatal("UpdateMDP: ", err)
	}
}
