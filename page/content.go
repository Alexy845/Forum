package page

import (
	"html/template"
	"log"
	"net/http"
)

func Content(w http.ResponseWriter, r *http.Request) {
	tmp := template.Must(template.ParseFiles("templates/html/content.html"))

	err := tmp.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}
