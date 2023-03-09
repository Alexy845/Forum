package page

import (
	"forum/requetes_sql"
	"forum/structs"
	uuid "github.com/satori/go.uuid"
	"log"
	"net/http"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		cookie, err := r.Cookie("session")
		if err == nil {
			id := uuid.Must(uuid.FromString(cookie.Value))
			structs.Datas.User = requetes_sql.GetUser(id)
			structs.Datas.Connected = true
			if r.FormValue("Type") == "comment" {
				idPost, err := uuid.FromString(r.FormValue("id-post"))
				if err != nil {
					log.Fatal(err)
				}
				structs.Datas.Post = requetes_sql.GetPost(idPost)
				idComment, err := uuid.FromString(r.FormValue("delete-comment"))
				if err != nil {
					log.Fatal(err)
				}
				requetes_sql.DeleteComment(idComment)
				if structs.Datas.Comment.Auteur.Id != id {
					http.Redirect(w, r, "/content?id="+structs.Datas.Post.Id.String(), http.StatusSeeOther)
				}
			} else if r.FormValue("Type") == "post" {
				idPost, err := uuid.FromString(r.FormValue("id-post"))
				if err != nil {
					log.Fatal(err)
				}
				if structs.Datas.Post.Auteur.Id == id {
					requetes_sql.DeletePost(idPost)
					http.Redirect(w, r, "/", http.StatusSeeOther)
				}
				http.Redirect(w, r, "/content?id="+structs.Datas.Post.Id.String(), http.StatusSeeOther)
			} else {
				http.Redirect(w, r, "/", http.StatusSeeOther)
			}
		} else {
			structs.Datas.Connected = false
			structs.Datas.Connected = false
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
	} else if r.Method == "GET" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
