package forum

import (
	"fmt"
	page "forum/page"
	"net/http"
)

func main() {
	http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("./templates"))))
	http.Handle("/dist/", http.StripPrefix("/dist/", http.FileServer(http.Dir("./dist"))))
	http.HandleFunc("/", page.Index)
	port := "80"
	fmt.Println("Startup Server on port " + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		return
	}
}
