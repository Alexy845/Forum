package page

import (
	"fmt"
	"forum/requetes_sql"
	"html/template"
	"log"
	"net/http"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/html/signup.html"))

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
			for i := 0; i < requetes_sql.GetTailleUser(); i++ {
				if requetes_sql.Getemail(i) == email {
					fmt.Println("Cette adresse email est déjà utilisée")
					break
				} else if requetes_sql.GetUser(i).Username == username {
					fmt.Println("Ce pseudo est déjà utilisé")
					break
				} else {
					requetes_sql.AddUser(username, firstname, lastname, password, email)
				}
			}
		}
	}
	err := tmpl.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}
