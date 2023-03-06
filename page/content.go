package page

import (
	"forum/requetes_sql"
	"forum/structs"
	uuid "github.com/satori/go.uuid"
	"html/template"
	"log"
	"net/http"
)

func Content(w http.ResponseWriter, r *http.Request) {
	tmp := template.Must(template.ParseFiles("templates/html/content.html"))
	cookie, err := r.Cookie("session")
	if err == nil {
		id := uuid.Must(uuid.FromString(cookie.Value))
		structs.Datas.User = requetes_sql.GetUser(id)
		structs.Datas.Connected = true
	}
	idPost, err := uuid.FromString(r.URL.Query().Get("id"))
	if err != nil {
		log.Fatal(err)
	}
	structs.Datas.Post = requetes_sql.GetPost(idPost)
	structs.Datas.Comments = requetes_sql.GetAllComment(idPost)
	structs.Datas.Post.Likes = requetes_sql.CountLike(idPost)
	err = tmp.Execute(w, structs.Datas)
	if err != nil {
		log.Fatal(err)
	}
}
