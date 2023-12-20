package main

import (
	"fmt"

	"github.com/6oof/miniweb-base/app"
	"github.com/6oof/miniweb-base/app/helpers"
	db "github.com/6oof/miniweb-base/database"
)

func main() {
	helpers.LoadEnv()
	appPort := helpers.EnvOrPanic("PORT")
	db.InitDB("hello.db")
	app.MbinServe(fmt.Sprintf(":%s", appPort))
}
