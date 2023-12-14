package db

import (
	"database/sql"

	"github.com/6oof/miniweb-base/database/sqlc"
	_ "github.com/mattn/go-sqlite3"
)

var Queries *sqlc.Queries

func NewDB(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func InitDB(dataSourceName string) {
	db, err := NewDB(dataSourceName)
	if err != nil {
		panic(err)
	}

	Queries = sqlc.New(db)
}
