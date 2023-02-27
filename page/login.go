package page

import (
	"html/template"
	"net/http"

)
func LoginFunc(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/html/login.html"))

	switch r.Method {
	case "GET":
	case "POST":
		username := r.FormValue("username")
		password := r.FormValue("password")
		if password != "" {
			if username != ""{
				if GetUserData(username).Username != "" && GetUserData(username).Password == password {
					CurrentUser = GetUserData(username)
				}
			}
		}
	}
	tmpl.Execute(w, nil)
}

func GetUserData(username string) User{
	users := _GetUsers()

	for i, user := range users {
		if user.Username == username {	
			return users[i]
		}
	}
	return User{Username:""}
}

