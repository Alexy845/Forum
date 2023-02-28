package requetes_sql

import (
	"database/sql"
	"log"
)

func GetTailleUser() int {
	req, err := DB.Query("SELECT count(*) FROM User")
	if err != nil {
		log.Fatal(err)
	}

	req.Next()
	taille := 0
	err = req.Scan(&taille)
	defer func(req *sql.Rows) {
		err := req.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(req)
	return taille
}
