package page

import (
	_const "forum/const"
	"forum/requetes_sql"
	"forum/structs"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func Search(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles(filepath.Join(_const.HtmlDir, "/search.html"))
	if err != nil {
		log.Fatal(err)
	}
	structs.Datas.Result = false
	input := r.URL.Query().Get("search")
	structs.Datas.Posts = requetes_sql.SearchPosts(input)
	for _, i := range requetes_sql.SearchUser(input) {
		for _, j := range requetes_sql.GetPostByUser(i.Id) {
			structs.Datas.Posts = append(structs.Datas.Posts, j)
		}
	}
	for _, i := range requetes_sql.SearchPostByCategory(input) {
		structs.Datas.Posts = append(structs.Datas.Posts, i)
	}
	if len(structs.Datas.Posts) > 0 {
		structs.Datas.Result = true
	}
	err = tmp.Execute(w, structs.Datas)
	if err != nil {
		log.Fatal(err)
	}
}
