package requetes_sql

import (
	"database/sql"
	"log"
)

func GetPseudo(id int) string {
	req, err := DB.Query("SELECT Pseudo FROM User")
	if err != nil {
		log.Fatal(err)
	}

	pseudo := ""
	err = req.Scan(&pseudo)
	defer func(req *sql.Rows) {
		err := req.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(req)
	req.Next()
	return pseudo
}
