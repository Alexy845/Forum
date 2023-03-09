package requetes_sql

import (
	uuid "github.com/satori/go.uuid"
	"log"
)

func DeleteComment(id uuid.UUID) {
	_, err := DB.Exec("DELETE FROM Comments WHERE Id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
}
