package page

import (
	"forum/requetes_sql"
	"forum/structs"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

func RAddComment(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err == nil {
		id := uuid.Must(uuid.FromString(cookie.Value))
		structs.Datas.User = requetes_sql.GetUser(id)
		structs.Datas.Connected = true
		requetes_sql.AddComment(structs.Datas.User.Id, uuid.FromStringOrNil(r.URL.Query().Get("PostID")), r.FormValue("Content"))
	}
	http.Redirect(w, r, "/content?id="+r.URL.Query().Get("id"), http.StatusSeeOther)
}
