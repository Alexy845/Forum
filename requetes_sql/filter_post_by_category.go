package requetes_sql

import (
	"database/sql"
	"forum/structs"
	uuid "github.com/satori/go.uuid"
	"log"
)

func FilterPostByCategory(category string, post string, liked string) []structs.Post {
	query_category := ""
	if category == "all" || category == "" {
		query_category = ""
	} else {
		query_category = " join Category on Posts.Category = Category.Id WHERE Category.Id = ?"
	}
	query_order := ""
	if post == "all" && liked == "all" {
		query_order = ""
	} else {
		query_order = " ORDER BY "
	}
	query_post := ""
	if post == "last" {
		query_post = "Posts.Date DESC"
	} else if post == "old" {
		query_post = "Posts.Date ASC"
	}
	query_like := ""
	if query_post != "" {
		query_like = ", "
	}
	if liked == "liked" {
		query_like += "Posts.Likes ASC"
	} else if liked == "not_liked" {
		query_like += "Posts.Dislikes ASC"
	}
	req, err := DB.Query("SELECT Posts.Id from Posts "+query_category+query_order+query_post+query_like, category)
	if err != nil {
		log.Fatal(err)
	}
	var posts []structs.Post
	for req.Next() {
		id := uuid.UUID{}
		err = req.Scan(&id)
		if err != nil {
			log.Fatal(err)
		}
		if err != nil {
			log.Fatal(err)
		}
		posts = append(posts, GetPost(id))
	}
	defer func(req *sql.Rows) {
		err := req.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(req)
	return posts
}
