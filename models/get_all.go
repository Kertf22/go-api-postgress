package models

import "go-api-postgress/db"

func GetAll() (todos []Todo, err error) {

	conn, err := db.OpenConnection()

	if err != nil {
		return todos, err
	}

	defer conn.Close()

	rows, err := conn.Query("SELECT * FROM todos")

	if err != nil {
		return todos, err
	}

	for rows.Next() {
		var todo Todo

		err = rows.Scan(&todo.ID, &todo.Title, &todo.Descr, &todo.Done)

		if err != nil {
			continue
		}

		todos = append(todos, todo)
	}

	return todos, err
}
