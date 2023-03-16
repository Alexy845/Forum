package requetes_sql

import (
	"database/sql"
	uuid "github.com/satori/go.uuid"
	"log"
)

func CountDislike(id uuid.UUID) int {
	req, err := DB.Query("SELECT COUNT(*) FROM Dislike WHERE Post = ?", id)
	if err != nil {
		log.Fatal()
	}
	i := 0
	for req.Next() {
		err = req.Scan(&i)
		if err != nil {
			log.Fatal(err)
		}
	}
	defer func(req *sql.Rows) {
		err := req.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(req)
	return i
}
