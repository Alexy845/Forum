package page

import (
	"fmt"
	"forum/cookies"
	"forum/requetes_sql"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

func Logining(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("session")
	if err == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	switch r.Method {
	case "GET":
		http.Redirect(w, r, "/", http.StatusSeeOther)
	case "POST":
		username := r.FormValue("username")
		password := r.FormValue("password")
		if password != "" && username != "" {
			id := requetes_sql.Verifuser(username, password)
			if id != uuid.Nil {
				http.SetCookie(w, cookies.CreateCookie(id))
				http.Redirect(w, r, "/", http.StatusSeeOther)
			} else {
				fmt.Println("Erreur de connexion")
			}
		}
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}
