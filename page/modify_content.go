package page

import (
	"forum/requetes_sql"
	"forum/structs"
	uuid "github.com/satori/go.uuid"
	"html/template"
	"log"
	"net/http"
)

func ModifyContent(w http.ResponseWriter, r *http.Request) {
	tmp := template.Must(template.ParseFiles("templates/html/modify_content.html"))
	cookie, err := r.Cookie("session")
	if err == nil {
		id := uuid.Must(uuid.FromString(cookie.Value))
		structs.Datas.User = requetes_sql.GetUser(id)
		structs.Datas.Connected = true
		if r.Method == "POST" {
			pseudo := r.FormValue("Pseudo")
			email := r.FormValue("Email")
			prenom := r.FormValue("Prenom")
			nom := r.FormValue("Nom")
			if pseudo == "" {
				pseudo = structs.Datas.User.Username
			}
			if email == "" {
				email = structs.Datas.User.Email
			}
			if prenom == "" {
				prenom = structs.Datas.User.FirstName
			}
			if nom == "" {
				nom = structs.Datas.User.LastName
			}
			requetes_sql.UpdateUser(id, pseudo, email, prenom, nom)
			http.Redirect(w, r, "/profile", http.StatusSeeOther)
		}
	} else {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	err = tmp.Execute(w, structs.Datas)
	if err != nil {
		log.Fatal(err)
	}
}
