package main

import (
	"fmt"
	"log"
	"net/http"
)

func headerHandler(writer http.ResponseWriter, request *http.Request) {
	h := request.Header
	fmt.Print(h)
	println()
	println("--------")
	println(h.Get("User-Agent"))

}
func main() {
	http.HandleFunc("/login", headerHandler)
	err := http.ListenAndServe("localhost:9977", nil)
	log.Fatal(err)
}
