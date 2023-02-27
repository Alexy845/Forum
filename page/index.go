package page

import (
	"forum/requetes_sql"
	"html/template"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/html/index.html")

	requetes_sql.Getall()

	err := t.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}
