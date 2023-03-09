package page

import (
	"forum/requetes_sql"
	"forum/structs"
	uuid "github.com/satori/go.uuid"
	"log"
	"net/http"
)

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		idPost, err := uuid.FromString(r.FormValue("id-post"))
		if err != nil {
			log.Fatal(err)
		}
		cookie, err := r.Cookie("session")
		if err == nil {
			id := uuid.Must(uuid.FromString(cookie.Value))
			structs.Datas.User = requetes_sql.GetUser(id)
			structs.Datas.Connected = true
			if r.FormValue("Type") == "comment" {
				idComment, err := uuid.FromString(r.FormValue("id"))
				if err != nil {
					log.Fatal(err)
				}
				requetes_sql.UpdateComment(idComment, r.FormValue("content"))
			} else if r.FormValue("Type") == "post" {
				requetes_sql.UpdatePost(idPost, r.FormValue("title"), r.FormValue("content"))
			}
		} else {
			structs.Datas.User = structs.User{}
			structs.Datas.Connected = false
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
		http.Redirect(w, r, "/content?id="+idPost.String(), http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther) // Modifier pour erreur
	}
}
