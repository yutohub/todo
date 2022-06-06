package model

import (
	"database/sql"

	_ "modernc.org/sqlite"

	"github.com/yutohub/todo/todo"
)

func CreateTable(db *sql.DB) error {
	const sql = `
	CREATE TABLE IF NOT EXISTS todo (
			id    INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			task  TEXT NOT NULL,
			done  INTEGER NOT NULL
	);`
	if _, err := db.Exec(sql); err != nil {
		return err
	}
	return nil
}

func ShowRecords(db *sql.DB) ([]todo.Todo, error) {
	todos := []todo.Todo{}
	rows, err := db.Query("SELECT * FROM todo")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var t todo.Todo
		if err := rows.Scan(&t.ID, &t.Task, &t.Done); err != nil {
			return nil, err
		}
		todos = append(todos, t)
	}
	return todos, nil
}

func InputRecord(t todo.Todo, db *sql.DB) error {
	const sql = "INSERT INTO todo(task, done) values (?,?)"
	_, err := db.Exec(sql, t.Task, t.Done)
	if err != nil {
		return err
	}
	return nil
}

func UpdateRecord(t todo.Todo, db *sql.DB) error {
	const sql = "UPDATE todo SET task = ?, done=? WHERE id = ?"
	_, err := db.Exec(sql, t.Task, t.Done, t.ID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteRecord(i int, db *sql.DB) error {
	const sql = "DELETE FROM todo WHERE id = ?"
	_, err := db.Exec(sql, i)
	if err != nil {
		return err
	}
	return nil
}
