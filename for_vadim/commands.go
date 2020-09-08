package main

import (
	"fmt"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func main () {
	dtdb, err := sql.Open("sqlite3", "testdb.db")
	if err != nil {
		panic(err)
	}
	defer dtdb.Close()
	query, err := dtdb.Exec("insert into list (Name, DataField) values ($1, $2)", "Vadim", "About me: Start to make this drisnia")
	if err != nil {
		panic(err)
	}

	fmt.Println(query.LastInsertId())
	fmt.Println(query.RowsAffected())


}