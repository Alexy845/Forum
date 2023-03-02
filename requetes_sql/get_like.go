package requetes_sql

import (
	"database/sql"
	"forum/structs"
	uuid "github.com/satori/go.uuid"
	"log"
)

func GetLike() structs.Like {
	req, err := DB.Query("SELECT Id, User, Post FROM Like")
	if err != nil {
		log.Fatal()
	}
	like := structs.Like{}
	for req.Next() {
		user := uuid.UUID{}
		post := uuid.UUID{}
		err = req.Scan(&like.Id, &user, &post)
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
