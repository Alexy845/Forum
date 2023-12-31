package page

import (
	"fmt"
	_const "forum/const"
	"forum/cookies"
	"forum/requetes_sql"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(filepath.Join(_const.HtmlDir, "signup.html")))
	_, err := r.Cookie("session")
	if err == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	switch r.Method {
	case "GET":
	case "POST":
		username := r.FormValue("username")
		password := r.FormValue("password")
		email := r.FormValue("email")
		firstname := r.FormValue("firstname")
		lastname := r.FormValue("lastname")
		confirmpassword := r.FormValue("confirmpassword")

		if password != confirmpassword {
			fmt.Println("Les mots de passe ne correspondent pas")
		} else {
			if requetes_sql.IfUserExist(email, username, firstname, lastname, password) {
				fmt.Println("L'utilisateur existe déjà")
			} else {
				http.SetCookie(w, cookies.CreateCookie(requetes_sql.AddUser(username, firstname, lastname, password, email)))
			}
		}
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}
