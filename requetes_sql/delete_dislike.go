package requetes_sql

import (
	uuid "github.com/satori/go.uuid"
	"log"
)

func DeleteDislike(id uuid.UUID) {
	_, err := DB.Exec("DELETE FROM Dislike WHERE Id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
}
