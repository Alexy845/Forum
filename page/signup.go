package page

import (
	"html/template"
	"net/http"

)
func SignUpFunc(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/html/signup.html"))
	tmpl.Execute(w, nil)
}