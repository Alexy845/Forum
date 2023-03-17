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

func Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/404", http.StatusSeeOther)
		return
	}
	t, _ := template.ParseFiles(filepath.Join(_const.HtmlDir, "/index.html"))
	cookie, err := r.Cookie("session")
	if err == nil {
		id := uuid.Must(uuid.FromString(cookie.Value))
		structs.Datas.User = requetes_sql.GetUser(id)
		structs.Datas.Connected = true
	} else {
		structs.Datas.User = structs.User{}
		structs.Datas.Connected = false
	}
	if r.Method == "POST" {
		category := r.FormValue("category")
		posts := r.FormValue("posts")
		liked := r.FormValue("liked")
		structs.Datas.Posts = requetes_sql.FilterPostByCategory(category, posts, liked)
	} else {
		structs.Datas.Posts = requetes_sql.GetAllPosts()
	}
	structs.Datas.Categories = requetes_sql.GetAllCategory()
	err = t.Execute(w, structs.Datas)
	if err != nil {
		log.Fatal(err)
	}
}
