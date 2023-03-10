package page

import (
	"fmt"
	_const "forum/const"
	"forum/requetes_sql"
	"forum/structs"
	uuid "github.com/satori/go.uuid"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func Update(w http.ResponseWriter, r *http.Request) {
	tmp := template.Must(template.ParseFiles(filepath.Join(_const.HtmlDir, "update.html")))
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
				idComment, err := uuid.FromString(r.FormValue("comment-update"))
				if err != nil {
					log.Fatal(err)
				}
				structs.Datas.Comment = requetes_sql.GetComment(idComment)
				if structs.Datas.Comment.Auteur.Id != id {
					http.Redirect(w, r, "/", http.StatusSeeOther) // TODO : Redirect to 404 Page
				} else {
					structs.Datas.EditPost = false
				}
			} else if r.FormValue("Type") == "post" {
				idPost, err := uuid.FromString(r.FormValue("id-post"))
				if err != nil {
					log.Fatal(err)
				}
				structs.Datas.Post = requetes_sql.GetPost(idPost)
				if structs.Datas.Post.Auteur.Id != id {
					http.Redirect(w, r, "/content?id="+structs.Datas.Post.Id.String(), http.StatusSeeOther)
				} else {
					structs.Datas.EditPost = true
				}
			} else {
				fmt.Println("ERROR")
			}
		} else {
			structs.Datas.User = structs.User{}
			structs.Datas.Connected = false
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
	} else if r.Method == "GET" {
		http.Redirect(w, r, "/", http.StatusSeeOther) // TODO : Redirect to 404 Page
	}
	err := tmp.Execute(w, structs.Datas)
	if err != nil {
		log.Fatal(err)
	}
}
