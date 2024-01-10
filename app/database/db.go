package db

import (
	"database/sql"

	"github.com/6oof/miniweb-base/app/database/sqlc"
	_ "modernc.org/sqlite"
)

// Queries represents the compiled SQL queries for interacting with the database.
var Queries *sqlc.Queries

// NewDB creates a new SQLite database connection with the specified data source name.
// It also applies additional settings for better performance.
func NewDB(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", dataSourceName) // Updated driver name
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

// InitDB initializes the database by creating a new SQLite database connection and
// instantiating the compiled SQL queries for the application.
func InitDB(dataSourceName string) {
	db, err := NewDB(dataSourceName)
	if err != nil {
		panic(err)
	}

	// Ensure that the database is properly closed when the application exits
	defer db.Close()

	// Create a new instance of the compiled SQL queries for the application
	Queries = sqlc.New(db)
}
