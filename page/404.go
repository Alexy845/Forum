package page

import (
	_const "forum/const"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func Page404(w http.ResponseWriter, r *http.Request) {
	tmp := template.Must(template.ParseFiles(filepath.Join(_const.HtmlDir, "404.html")))
	err := tmp.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}
