package page

import (
	"forum/structs"
	"net/http"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	if !structs.Datas.Connected {
		structs.Datas.User = structs.User{}
		structs.Datas.Connected = false
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		cookie, err := r.Cookie("session")
		if err != nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		cookie.MaxAge = -1
		http.SetCookie(w, cookie)
		structs.Datas.Connected = false
		structs.Datas.User = structs.User{}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
