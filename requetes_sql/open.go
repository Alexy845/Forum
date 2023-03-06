package requetes_sql

import (
	"database/sql"
	_const "forum/const"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func Open() *sql.DB {
	var sqliteDatabase, err = sql.Open("sqlite3", _const.Database)
	if err != nil {
		log.Fatal("DB : ", err)
	}
	return sqliteDatabase
}
