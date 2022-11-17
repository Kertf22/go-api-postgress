package models

import "go-api-postgress/db"

func Update(id int64, todo Todo) (int64, error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return 0, err
	}

	defer conn.Close()

	res, err := conn.Exec(
		`UPDATE todos set title=$1, description=$2, done=$3 where id=$4`,
		todo.Title, todo.Descr, todo.Done, id)

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
