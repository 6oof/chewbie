package main

import (
	"fmt"

	"github.com/6oof/miniweb-base/app"
	"github.com/6oof/miniweb-base/app/helpers"
	db "github.com/6oof/miniweb-base/database"
)

func main() {
	// Load environment variables from the .env file
	helpers.LoadEnv()

	// Get the application port from the environment variables or panic if not set
	appPort := helpers.EnvOrPanic("PORT")

	// Initialize the database with the specified file
	db.InitDB("mw.db")

	// Start the MiniWeb server
	app.MbinServe(fmt.Sprintf(":%s", appPort))
}
