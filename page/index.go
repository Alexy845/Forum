package page

import (
	"fmt"
	"forum/requetes_sql"
	uuid "github.com/satori/go.uuid"
	"html/template"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/html/index.html")
	cookie, err := r.Cookie("session")
	if err != nil {
		return
	}

	id := uuid.Must(uuid.FromString(cookie.Value))
	fmt.Println(requetes_sql.GetUser(id))
	err = t.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}
