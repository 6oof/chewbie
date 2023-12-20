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

	// Enable Write-Ahead Logging (WAL) mode
	_, err = db.Exec("PRAGMA journal_mode=WAL;")
	if err != nil {
		return nil, err
	}

	// Additional settings for better performance
	_, err = db.Exec("PRAGMA synchronous=NORMAL;")
	if err != nil {
		return nil, err
	}

	// You can experiment with other PRAGMA settings based on your application's requirements.

	return db, nil
}

func InitDB(dataSourceName string) {
	db, err := NewDB(dataSourceName)
	if err != nil {
		panic(err)
	}

	// Ensure that the database is properly closed when the application exits
	defer db.Close()

	Queries = sqlc.New(db)

	// Perform any additional initialization steps here, such as creating tables or applying migrations.
	// Make sure to handle any errors that may occur during these steps.
}
