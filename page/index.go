package page

import (
	"html/template"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/html/index.html")
	err := t.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}
