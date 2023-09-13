package db

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pablouser1/NotesManager/constants/files"
)

var conn *sql.DB

func Open(basePath string) {
	path := filepath.Join(basePath, "data.db")
	var err error
	conn, err = sql.Open("sqlite3", path)
	if err != nil {
		fmt.Println("Could not open db")
		return
	}

	// Create tables if database does not exist
	if _, err := os.Stat(path); err != nil {
		conn.Exec(files.DB_STRUCT)
	}
}

func Close() {
	conn.Close()
}
