package requetes_sql

import (
	"fmt"
	uuid2 "forum/uuid"
	"log"
)

func AddCategory(name string) {
	_, err := DB.Exec("INSERT INTO Category (Id, Name) VALUES (?, ?)", uuid2.CreateUUID(), name)
	if err != nil {
		fmt.Println("Error while adding category")
		log.Fatal(err)
	}
}
