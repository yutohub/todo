package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/yutohub/todo/model"
	"github.com/yutohub/todo/todo"
)

func TodoHandler(w http.ResponseWriter, r *http.Request) {
	log.Print(r.Method)
	model.CreateTable(model.Db)
	switch r.Method {
	case http.MethodGet:
		todos, err := model.ShowRecords(model.Db)
		if err != nil {
			log.Print(err)
		}
		todoResponse, err := json.MarshalIndent(todos, "", "\t")
		if err != nil {
			log.Print(err)
		}
		w.Write(todoResponse)
	case http.MethodPost:
		var todoRequest todo.Todo
		body := make([]byte, r.ContentLength)
		r.Body.Read(body)
		json.Unmarshal(body, &todoRequest)
		err := model.InputRecord(todoRequest, model.Db)
		if err != nil {
			log.Print(err)
		}
	case http.MethodPut:
		var todoRequest todo.Todo
		body := make([]byte, r.ContentLength)
		r.Body.Read(body)
		json.Unmarshal(body, &todoRequest)
		err := model.UpdateRecord(todoRequest, model.Db)
		if err != nil {
			log.Print(err)
		}
	case http.MethodDelete:
		todoId, _ := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/todo/"))
		err := model.DeleteRecord(todoId, model.Db)
		if err != nil {
			log.Print(err)
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, "Method not allowed.\n")
	}
}
