package page

import (
	"fmt"
	_const "forum/const"
	"forum/requetes_sql"
	"forum/structs"
	uuid "github.com/satori/go.uuid"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func Index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles(filepath.Join(_const.HtmlDir, "/index.html"))
	cookie, err := r.Cookie("session")
	if err == nil {
		id := uuid.Must(uuid.FromString(cookie.Value))
		structs.Datas.User = requetes_sql.GetUser(id)
		structs.Datas.Connected = true
	}
	structs.Datas.Posts = requetes_sql.GetAllPosts()
	fmt.Println(filepath.Join(_const.HtmlDir, "/index.html"))
	err = t.Execute(w, structs.Datas)
	if err != nil {
		log.Fatal(err)
	}
}
