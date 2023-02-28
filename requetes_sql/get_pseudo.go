package requetes_sql

import (
	"database/sql"
	"fmt"
	"log"
)

func GetPseudo(id int) string {
	req, err := DB.Query("SELECT Pseudo FROM User")
	if err != nil {
		fmt.Println("GetPseudo1")
		log.Fatal(err)
	}

	pseudo := ""
	err = req.Scan(&pseudo)
	defer func(req *sql.Rows) {
		err := req.Close()
		if err != nil {
			fmt.Println("GetPseudo2")
			log.Fatal(err)
		}
	}(req)
	req.Next()
	return pseudo
}
