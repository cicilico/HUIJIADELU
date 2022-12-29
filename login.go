package main

import (
	"html/template"
	"log"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("../gohome/login.html")
		t.Execute(w, nil)

	}

}
func main() {
	http.HandleFunc("/login", login)
	err := http.ListenAndServe("localhost:9977", nil)
	log.Fatal(err)
}
