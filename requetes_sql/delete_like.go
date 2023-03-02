package requetes_sql

import (
	uuid "github.com/satori/go.uuid"
	"log"
)

func DeleteLike(id uuid.UUID) {
	_, err := DB.Exec("DELETE FROM Like WHERE Id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
}
