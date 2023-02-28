package page

import (
	"forum/requetes_sql"
	uuid "github.com/satori/go.uuid"
	"log"
	"net/http"
)

func Writing(w http.ResponseWriter, r *http.Request) {
	titre := r.FormValue("Title")
	contenu := r.FormValue("Content")
	cookie, err := r.Cookie("session")
	if err != nil {
		log.Fatal(err)
	}
	id := uuid.Must(uuid.FromString(cookie.Value))
	requetes_sql.AddPost(id, titre, contenu)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
