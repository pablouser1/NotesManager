package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var conn *sql.DB

func Open(basePath string) {
	var err error
	conn, err = sql.Open("sqlite3", basePath+"/data.db")
	if err != nil {
		fmt.Println("Could not open db")
	}
}

func Close() {
	conn.Close()
}
