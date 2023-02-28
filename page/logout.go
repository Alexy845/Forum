package page

import (
	"forum/structs"
	"net/http"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	cookie.MaxAge = -1
	http.SetCookie(w, cookie)
	structs.Datas.Connected = false
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
