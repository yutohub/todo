package model

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("sqlite", "model/todo.db")
	if err != nil {
		log.Print(err)
	}
}
