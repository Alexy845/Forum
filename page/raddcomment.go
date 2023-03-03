package page

import (
	"fmt"
	"forum/requetes_sql"
	"forum/structs"
	uuid "github.com/satori/go.uuid"
	"log"
	"net/http"
)

func RAddComment(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err == nil {
		id := uuid.Must(uuid.FromString(cookie.Value))
		structs.Datas.User = requetes_sql.GetUser(id)
		structs.Datas.Connected = true
		postid, err := uuid.FromString(r.FormValue("post"))
		if err != nil {
			log.Fatal(err)
		}
		requetes_sql.AddComment(structs.Datas.User.Id, postid, r.FormValue("Comment"))
	}
	postid, err := uuid.FromString(r.FormValue("post"))
	if err != nil {
		log.Fatal(err)
	}
	url := "/content?id=" + postid.String()
	fmt.Println(url)
	http.Redirect(w, r, url, http.StatusSeeOther)
}
