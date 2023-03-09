package page

import (
	_const "forum/const"
	"forum/structs"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func LoginPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(filepath.Join(_const.HtmlDir, "login.html")))
	_, err := r.Cookie("session")
	if err == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		structs.Datas.User = structs.User{}
		structs.Datas.Connected = false
	}
	err = tmpl.Execute(w, structs.Datas)
	if err != nil {
		log.Fatal(err)
	}
}
