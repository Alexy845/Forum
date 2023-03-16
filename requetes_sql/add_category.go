package requetes_sql

import (
	uuid2 "forum/uuid"
	"log"
)

func AddCategory(name string) {
	_, err := DB.Exec("INSERT INTO Category (Id, Name) VALUES (?, ?)", uuid2.CreateUUID(), name)
	if err != nil {
		log.Fatal(err)
	}
}
