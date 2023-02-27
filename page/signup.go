package page

import (
	"fmt"
	"forum/requetes_sql"
	"html/template"
	"net/http"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/html/signup.html"))

	switch r.Method {
	case "GET":
	case "POST":
		username := r.FormValue("username")
		password := r.FormValue("password")

		if password != "" && username != "" {
			id := requetes_sql.Verifuser(username, password)
			if id != 0 {
				CurrentUser := requetes_sql.GetUser(id)
				fmt.Println(CurrentUser)
			} else {
				fmt.Println("Erreur de connexion")
			}
		}
	}
	tmpl.Execute(w, nil)
}
