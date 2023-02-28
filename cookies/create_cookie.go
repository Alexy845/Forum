package cookies

import (
	uuid "github.com/satori/go.uuid"
	"net/http"
	"time"
)

func CreateCookie(id uuid.UUID) *http.Cookie {
	userIdCookie := &http.Cookie{
		Name:       "session",
		Value:      id.String(),
		Path:       "",
		Domain:     "",
		Expires:    time.Time{},
		RawExpires: "",
		MaxAge:     99999999999,
		Secure:     false,
		HttpOnly:   false,
		SameSite:   0,
		Raw:        "",
		Unparsed:   []string{},
	}
	return userIdCookie
}
