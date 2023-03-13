package page

import (
	_const "forum/const"
	"forum/requetes_sql"
	"forum/structs"
	uuid "github.com/satori/go.uuid"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func Content(w http.ResponseWriter, r *http.Request) {
	tmp := template.Must(template.ParseFiles(filepath.Join(_const.HtmlDir, "content.html")))
	cookie, err := r.Cookie("session")
	if err == nil {
		id := uuid.Must(uuid.FromString(cookie.Value))
		structs.Datas.User = requetes_sql.GetUser(id)
		structs.Datas.Connected = true
	} else {
		structs.Datas.Connected = false
		structs.Datas.Connected = false
	}
	idPost, err := uuid.FromString(r.URL.Query().Get("id"))
	if err != nil {
		http.Redirect(w, r, "/404", http.StatusSeeOther)
	}
	structs.Datas.Post = requetes_sql.GetPost(idPost)
	structs.Datas.Comments = requetes_sql.GetAllComment(idPost)
	structs.Datas.Post.Likes = requetes_sql.CountLike(idPost)
	err = tmp.Execute(w, structs.Datas)
	if err != nil {
		log.Fatal(err)
	}
}
