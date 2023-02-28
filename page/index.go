package page

import (
	"forum/requetes_sql"
	"forum/structs"
	uuid "github.com/satori/go.uuid"
	"html/template"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/html/index.html")
	cookie, err := r.Cookie("session")
	if err != nil {
		log.Fatal(err)
	}

	id := uuid.Must(uuid.FromString(cookie.Value))
	structs.Datas.Posts = requetes_sql.GetAllPosts()
	structs.Datas.User = requetes_sql.GetUser(id)
	err = t.Execute(w, structs.Datas)
	if err != nil {
		log.Fatal(err)
	}
}
