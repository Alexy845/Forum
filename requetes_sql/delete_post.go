package requetes_sql

import (
	uuid "github.com/satori/go.uuid"
	"log"
)

func DeletePost(id uuid.UUID) {
	_, err := DB.Exec("DELETE FROM Posts WHERE Id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
}
