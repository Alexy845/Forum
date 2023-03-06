package main

import (
	"fmt"
	page "forum/page"
	"log"
	"net/http"
)

func main() {
	http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("./templates"))))
	http.Handle("/dist/", http.StripPrefix("/dist/", http.FileServer(http.Dir("./dist"))))
	http.HandleFunc("/", page.Index)
	http.HandleFunc("/login", page.LoginPage)
	http.HandleFunc("/logining", page.Logining)
	http.HandleFunc("/signup", page.Signup)
	http.HandleFunc("/write", page.Write)
	http.HandleFunc("/writing", page.Writing)
	http.HandleFunc("/logout", page.Logout)
	http.HandleFunc("/profile", page.Profile)
	http.HandleFunc("/content", page.Content)
	http.HandleFunc("/addcomment", page.RAddComment)
	http.HandleFunc("/modify", page.Modify)
	port := "8080"
	fmt.Println("Startup Server on port " + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
