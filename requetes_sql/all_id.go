package requetes_sql

import (
	"database/sql"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"log"
)

func AllId() []uuid.UUID {
	req, err := DB.Query("SELECT id FROM User")
	if err != nil {
		fmt.Println("GetUser1")
		log.Fatal()
	}
	ids := []uuid.UUID{}
	id := uuid.UUID{}
	for req.Next() {
		err = req.Scan(&id)
		if err != nil {
			fmt.Println("GetUser2")
			log.Fatal(err)
		}
		ids = append(ids, id)
	}
	defer func(req *sql.Rows) {
		err := req.Close()
		if err != nil {
			fmt.Println("GetUser3")
			log.Fatal(err)
		}
	}(req)
	return ids
}
