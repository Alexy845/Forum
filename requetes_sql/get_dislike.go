package requetes_sql

import (
	"database/sql"
	"forum/structs"
	uuid "github.com/satori/go.uuid"
	"log"
)

func GetDislike(user uuid.UUID, post uuid.UUID) structs.Dislike {
	req, err := DB.Query("SELECT Id FROM Dislike WHERE User = ? and Post = ?", user, post)
	if err != nil {
		log.Fatal()
	}
	dislike := structs.Dislike{}
	for req.Next() {
		err = req.Scan(&dislike.Id)
		if err != nil {
			log.Fatal(err)
		}
		dislike.User = GetUser(user)
		dislike.Post = GetPost(post)
	}
	defer func(req *sql.Rows) {
		err := req.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(req)
	return dislike
}
