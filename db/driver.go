package db

import (
	"database/sql"
	"fmt"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

var conn *sql.DB

func Open(basePath string) {
	var err error
	conn, err = sql.Open("sqlite3", filepath.Join(basePath, "data.db"))
	if err != nil {
		fmt.Println("Could not open db")
	}
}

func Close() {
	conn.Close()
}
