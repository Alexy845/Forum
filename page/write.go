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

func Write(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles(filepath.Join(_const.HtmlDir, "writePost.html"))
	if err != nil {
		log.Fatal(err)
	}

	cookie, err := r.Cookie("session")
	if err != nil {
		structs.Datas.User = structs.User{}
		structs.Datas.Connected = false
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {
		id := uuid.Must(uuid.FromString(cookie.Value))
		structs.Datas.User = requetes_sql.GetUser(id)
	}
	err = tmp.Execute(w, structs.Datas)
	if err != nil {
		log.Fatal(err)
	}
}
