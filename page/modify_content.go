package page

import (
	_const "forum/const"
	"forum/requetes_sql"
	"forum/structs"
	uuid "github.com/satori/go.uuid"
	"html/template"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func ModifyContent(w http.ResponseWriter, r *http.Request) {
	tmp := template.Must(template.ParseFiles(filepath.Join(_const.HtmlDir, "modify_content.html")))
	cookie, err := r.Cookie("session")
	if err == nil {
		id := uuid.Must(uuid.FromString(cookie.Value))
		structs.Datas.User = requetes_sql.GetUser(id)
		structs.Datas.Connected = true
		if r.Method == "POST" {
			pseudo := r.FormValue("Pseudo")
			email := r.FormValue("Email")
			prenom := r.FormValue("Prenom")
			nom := r.FormValue("Nom")
			file, fileheader, _ := r.FormFile("avatar")
			avatar := structs.Datas.User.Avatar
			if file != nil {
				defer func(file multipart.File) {
					err := file.Close()
					if err != nil {
						log.Fatal(err)
					}
				}(file)
				_, err := os.Stat(_const.Root + "/templates/image/avatars/" + id.String())
				if err != nil {
					err := os.Mkdir(_const.Root+"/templates/image/avatars/"+id.String(), 0755)
					if err != nil {
						log.Fatal("Create folder : ", err)
					}
				} else {
					err := os.Remove(_const.Root + structs.Datas.User.Avatar)
					if err != nil {
						log.Fatal("Remove folder : ", err)
					}
				}
				avatar = "/templates/image/avatars/" + id.String() + "/" + fileheader.Filename
				out, err := os.Create(_const.Root + avatar)
				if err != nil {
					log.Fatal("Create file : ", err)
				}
				_, err = io.Copy(out, file)
				if err != nil {
					log.Fatal(err)
				}
			}
			if pseudo == "" {
				pseudo = structs.Datas.User.Username
			}
			if email == "" {
				email = structs.Datas.User.Email
			}
			if prenom == "" {
				prenom = structs.Datas.User.FirstName
			}
			if nom == "" {
				nom = structs.Datas.User.LastName
			}
			requetes_sql.UpdateUser(id, pseudo, email, prenom, nom, avatar)
			http.Redirect(w, r, "/profile", http.StatusSeeOther)
		}
	} else {
		structs.Datas.User = structs.User{}
		structs.Datas.Connected = false
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
	err = tmp.Execute(w, structs.Datas)
	if err != nil {
		log.Fatal(err)
	}
}
