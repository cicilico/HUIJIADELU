package main

import (
	"log"
	"net/http"
)

func writerHandler(writer http.ResponseWriter, request *http.Request) {
	htm := `<html>
    <head>
        <title>request 处理</title>
    </head>
    <body>
        <form action="login?userName=张三&password=123" method="post">
            用户名：<input type="text" name="userName1">
            密码：<input type="password" name="password1">
            <input type="submit" value="登录">
        </form>
    </body>
</html>`

	writer.WriteHeader(200)
	writer.Write([]byte(htm))
}
func main() {
	http.HandleFunc("/login", writerHandler)
	err := http.ListenAndServe("localhost:9977", nil)
	log.Fatal(err)
}
