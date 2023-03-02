package page

import (
	"forum/requetes_sql"
	"forum/structs"
	uuid "github.com/satori/go.uuid"
	"html/template"
	"net/http"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	tmp := template.Must(template.ParseFiles("templates/html/profile.html"))
	cookie, err := r.Cookie("session")
	if err == nil {
		id := uuid.Must(uuid.FromString(cookie.Value))
		structs.Datas.User = requetes_sql.GetUser(id)
		structs.Datas.Connected = true
	}
	err = tmp.Execute(w, structs.Datas)
	if err != nil {
		return
	}
}
