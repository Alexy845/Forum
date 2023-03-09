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
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
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
		} else {
			structs.Datas.User = structs.User{}
			structs.Datas.Connected = false
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
		postid, err := uuid.FromString(r.FormValue("post"))
		if err != nil {
			log.Fatal(err)
		}
		url := "/content?id=" + postid.String()
		fmt.Println(url)
		http.Redirect(w, r, url, http.StatusSeeOther)
	}
}
