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

func Category(w http.ResponseWriter, r *http.Request) {
	tmp := template.Must(template.ParseFiles(filepath.Join(_const.HtmlDir, "category.html")))
	cookie, err := r.Cookie("session")
	if err == nil {
		id := uuid.Must(uuid.FromString(cookie.Value))
		structs.Datas.User = requetes_sql.GetUser(id)
		structs.Datas.Connected = true
	} else {
		structs.Datas.Connected = false
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
	if r.Method == "POST" {
		requetes_sql.AddCategory(r.FormValue("category"))
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	err = tmp.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}
