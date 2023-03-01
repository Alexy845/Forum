package requetes_sql

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func Open() *sql.DB {
	var sqliteDatabase, _ = sql.Open("sqlite3", "./DB_Base.sqlite")
	return sqliteDatabase
}
