package repository

import (
	"database/sql"
)

func Connect() {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:13306)/user")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}
