package page

import (
	"forum/requetes_sql"
	"forum/structs"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

func Writing(w http.ResponseWriter, r *http.Request) {
	titre := r.FormValue("Title")
	contenu := r.FormValue("Content")
	category := r.FormValue("Category")
	cookie, err := r.Cookie("session")
	if err != nil {
		structs.Datas.User = structs.User{}
		structs.Datas.Connected = false
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {
		id := uuid.Must(uuid.FromString(cookie.Value))
		Thecategory := uuid.Must(uuid.FromString(category))
		requetes_sql.AddPost(id, titre, contenu, Thecategory)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
