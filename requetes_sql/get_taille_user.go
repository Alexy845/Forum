package requetes_sql

import (
	"database/sql"
	"fmt"
	"log"
)

func GetTailleUser() int {
	req, err := DB.Query("SELECT count(*) FROM User")
	if err != nil {
		fmt.Println("GetTailleUser1")
		log.Fatal(err)
	}

	req.Next()
	taille := 0
	err = req.Scan(&taille)
	defer func(req *sql.Rows) {
		err := req.Close()
		if err != nil {
			fmt.Println("GetTailleUser2")
			log.Fatal(err)
		}
	}(req)
	return taille
}
