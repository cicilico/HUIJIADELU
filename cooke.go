package main

import (
	"log"
	"net/http"
)

func getCookie(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("username")
	if err != nil {
		log.Fatal(err)
	}
	println(cookie.Value)
}

func setCookie(writer http.ResponseWriter, request *http.Request) {
	cookie := http.Cookie{Name: "username", Value: "Saddam"}
	http.SetCookie(writer, &cookie)

}

func main() {
	http.HandleFunc("/setCookie", setCookie)
	http.HandleFunc("/getCookie", getCookie)

	err := http.ListenAndServe("localhost:9977", nil)
	log.Fatal(err)

}
