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
	cookie, err := r.Cookie("session")
	if err != nil {
		structs.Datas.User = structs.User{}
		structs.Datas.Connected = false
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {
		id := uuid.Must(uuid.FromString(cookie.Value))
		requetes_sql.AddPost(id, titre, contenu)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
