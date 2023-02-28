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
	http.HandleFunc("/login", page.LoginFunc)
	http.HandleFunc("/signup", page.Signup)
	port := "8080"
	fmt.Println("Startup Server on port " + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
