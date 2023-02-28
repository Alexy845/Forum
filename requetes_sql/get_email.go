package requetes_sql

import (
	"database/sql"
	"log"
)

func Getemail(id int) string {
	req, err := DB.Query("SELECT email FROM User where Id = ?", id)
	if err != nil {
		log.Fatal(err)
	}

	req.Next()
	email := ""
	err = req.Scan(&email)
	defer func(req *sql.Rows) {
		err := req.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(req)
	return email
}
