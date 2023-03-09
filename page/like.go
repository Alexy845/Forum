package page

import (
	"forum/requetes_sql"
	"forum/structs"
	uuid "github.com/satori/go.uuid"
	"log"
	"net/http"
)

func Like(w http.ResponseWriter, r *http.Request) {
	id_post, err := uuid.FromString(r.FormValue("like"))
	if err != nil {
		log.Fatal(err)
	}
	cookie, err := r.Cookie("session")
	if err == nil {
		id := uuid.Must(uuid.FromString(cookie.Value))
		structs.Datas.User = requetes_sql.GetUser(id)
		structs.Datas.Connected = true
		if requetes_sql.GetLike(id, id_post) == (structs.Like{}) {
			requetes_sql.AddLike(id_post, structs.Datas.User.Id)
		} else {
			requetes_sql.DeleteLike(requetes_sql.GetLike(id, id_post).Id)
		}
	}
	http.Redirect(w, r, "/content?id="+id_post.String(), http.StatusSeeOther)
}