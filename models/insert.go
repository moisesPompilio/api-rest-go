package models

import (
	"github.com/moisesPompilio/api-rest-go/db"
)

func Insert(todo Todo) (id int64, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}

	defer conn.Close()

	err = conn.QueryRow("INSERT INTO todos (title, description, done) VALUES ($1, $2, $3) RETURNING id",
		todo.Title, todo.Description, todo.Done).Scan(&id)

	return
}
