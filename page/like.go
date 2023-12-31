package page

import (
	"forum/requetes_sql"
	"forum/structs"
	uuid "github.com/satori/go.uuid"
	"log"
	"net/http"
)

func Like(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/404", http.StatusSeeOther)
	} else {
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
				if requetes_sql.GetDislike(id, id_post) != (structs.Dislike{}) {
					requetes_sql.DeleteDislike(requetes_sql.GetDislike(id, id_post).Id)
					requetes_sql.UpdateDislike(id_post, requetes_sql.CountDislike(id_post))
				}
				requetes_sql.AddLike(id_post, structs.Datas.User.Id)
			} else {
				requetes_sql.DeleteLike(requetes_sql.GetLike(id, id_post).Id)
			}
			requetes_sql.UpdateLike(id_post, requetes_sql.CountLike(id_post))
			http.Redirect(w, r, "/content?id="+id_post.String(), http.StatusSeeOther)
		} else {
			structs.Datas.User = structs.User{}
			structs.Datas.Connected = false
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
	}
}
