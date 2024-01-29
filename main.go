package main

import (
	"fmt"

	"github.com/6oof/minweb/app"
	"github.com/6oof/minweb/app/helpers"
	// Uncomment line below to use the default database file (1/2)
	// db "github.com/6oof/minweb/database"
)

func main() {
	// Load environment variables from the .env file
	helpers.LoadEnv()

	// Int logger
	helpers.InitLogger("log.txt")

	// Get the application port from the environment variables or panic if not set
	appPort := helpers.Env("PORT", ":3003")

	// Initialize the database with the specified file
	// uncomment line line below to use the default database file (2/2)
	// db.InitDB("mw.db")
	// Follow instructions in the documentation to migrate and generate queries

	// Start the MiniWeb server
	app.MbinServe(fmt.Sprintf(":%s", appPort))
}
