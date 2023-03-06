package page

import (
	_const "forum/const"
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
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}
