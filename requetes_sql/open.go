package requetes_sql

import (
	"database/sql"
	_const "forum/const"
	_ "github.com/mattn/go-sqlite3"
)

func Open() *sql.DB {
	var sqliteDatabase, _ = sql.Open("sqlite3", _const.Database)
	return sqliteDatabase
}
