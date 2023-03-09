package page

import (
	"forum/hash"
	"forum/requetes_sql"
	"forum/structs"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

func Modify(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err == nil {
		id := uuid.Must(uuid.FromString(cookie.Value))
		structs.Datas.User = requetes_sql.GetUser(id)
		structs.Datas.Connected = true
		if r.Method == "POST" {
			if hash.CompareHash(structs.Datas.User.Password, r.FormValue("password_old")) {
				mdp := r.FormValue("password_new")
				if mdp != "" && mdp == r.FormValue("password_confirm") {
					requetes_sql.UpdateMDP(hash.Hash(mdp))
				}
			}
		}
	} else {
		structs.Datas.User = structs.User{}
		structs.Datas.Connected = false
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
	http.Redirect(w, r, "/profile", http.StatusSeeOther)
}
