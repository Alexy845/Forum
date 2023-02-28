package page

import (
	"html/template"
	"log"
	"net/http"
)

func LoginPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/html/login.html"))
	_, err := r.Cookie("session")
	if err == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}
