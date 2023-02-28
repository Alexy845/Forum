package page

import (
	"forum/requetes_sql"
	"forum/structs"
	uuid "github.com/satori/go.uuid"
	"html/template"
	"log"
	"net/http"
)

func Write(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("templates/html/writePost.html")
	if err != nil {
		log.Fatal(err)
	}

	cookie, err := r.Cookie("session")
	if err != nil {
		log.Fatal(err)
	}
	id := uuid.Must(uuid.FromString(cookie.Value))
	structs.Datas.User = requetes_sql.GetUser(id)
	err = tmp.Execute(w, structs.Datas)
	if err != nil {
		log.Fatal(err)
	}
}
