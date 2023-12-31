package requetes_sql

import (
	"database/sql"
	"forum/structs"
	uuid "github.com/satori/go.uuid"
	"log"
)

func GetLike(user uuid.UUID, post uuid.UUID) structs.Like {
	req, err := DB.Query("SELECT Id FROM Like WHERE User = ? and Post = ?", user, post)
	if err != nil {
		log.Fatal()
	}
	like := structs.Like{}
	for req.Next() {
		err = req.Scan(&like.Id)
		if err != nil {
			log.Fatal(err)
		}
		like.User = GetUser(user)
		like.Post = GetPost(post)
	}
	defer func(req *sql.Rows) {
		err := req.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(req)
	return like
}
