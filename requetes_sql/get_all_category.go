package requetes_sql

import (
	"database/sql"
	"forum/structs"
	"log"
)

func GetAllCategory() []structs.Category {
	req, err := DB.Query("SELECT * FROM Category")
	if err != nil {
		log.Fatal(err)
	}
	lst_category := []structs.Category{}
	for req.Next() {
		category := structs.Category{}
		err = req.Scan(&category.Id, &category.Name)
		if err != nil {
			log.Fatal(err)
		}
		lst_category = append(lst_category, category)
	}
	defer func(req *sql.Rows) {
		err := req.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(req)
	return lst_category
}
