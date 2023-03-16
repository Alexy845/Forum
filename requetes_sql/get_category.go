package requetes_sql

import (
	"database/sql"
	"forum/structs"
	uuid "github.com/satori/go.uuid"
	"log"
)

func GetCategory(id uuid.UUID) structs.Category {
	req, err := DB.Query("SELECT * FROM Category where Id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	category := structs.Category{}
	for req.Next() {
		err = req.Scan(&category.Id, &category.Name)
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
	return category
}
