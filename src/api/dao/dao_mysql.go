package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func StartMySQL(user, pass string) {

	db, err := sql.Open("mysql", "b06e9ae78a9434:f650ac03@@us-cdbr-east-03.cleardb.com/heroku_854bd42de5b62e1?reconnect=true")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	// See "Important settings" section.
	// perform a db.Query insert
	insert, err := db.Query("INSERT INTO users VALUES ( ?, ? )", user, pass)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()

}