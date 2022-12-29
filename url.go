package main

import (
	"html/template"
	"log"
	"net/http"
)

func loginHandler(writer http.ResponseWriter, request *http.Request) {
	method := request.Method
	println("请求方法", method)

	if request.Method == "GET" {
		t, _ := template.ParseFiles("login.html")
		t.Execute(writer, nil)

	} else {

		url := request.URL
		values := url.Query()
		username := values.Get("username")
		password := values.Get(" password")
		println(username, password)

		request.ParseForm()
		username1 := request.Form.Get("username1")
		password1 := request.Form.Get("password1")
		println(username1, password1)
	}

}
func main() {
	http.HandleFunc("/login", loginHandler)
	err := http.ListenAndServe("localhost:9977", nil)
	log.Fatal(err)

}
