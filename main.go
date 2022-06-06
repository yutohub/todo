package main

import (
	"net/http"

	"github.com/yutohub/todo/controller"
)

func main() {
	http.HandleFunc("/todo/", controller.TodoHandler)
	http.ListenAndServe(":8080", nil)
}
