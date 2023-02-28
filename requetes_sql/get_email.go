package requetes_sql

import (
	"database/sql"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"log"
)

func Getemail(id uuid.UUID) string {
	req, err := DB.Query("SELECT email FROM User where Id = ?", id)
	if err != nil {
		fmt.Println("Getemail1")
		log.Fatal(err)
	}

	req.Next()
	email := ""
	err = req.Scan(&email)
	defer func(req *sql.Rows) {
		err := req.Close()
		if err != nil {
			fmt.Println("Getemail2")
			log.Fatal(err)
		}
	}(req)
	return email
}
