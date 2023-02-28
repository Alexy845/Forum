package page

import (
	"fmt"
	"forum/cookies"
	"forum/requetes_sql"
	uuid "github.com/satori/go.uuid"
	"html/template"
	"log"
	"net/http"
)

func LoginFunc(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/html/login.html"))
	switch r.Method {
	case "GET":
	case "POST":
		username := r.FormValue("username")
		password := r.FormValue("password")
		if password != "" && username != "" {
			id := requetes_sql.Verifuser(username, password)
			if id != uuid.Nil {
				http.SetCookie(w, cookies.CreateCookie(id))
				CurrentUser := requetes_sql.GetUser(id)
				fmt.Println(CurrentUser)
			} else {
				fmt.Println("Erreur de connexion")
			}
		}
	}
	err := tmpl.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}
