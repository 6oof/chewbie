package main

import (
	"github.com/6oof/miniweb-base/app"
	db "github.com/6oof/miniweb-base/database"
)

func main() {

	db.InitDB("hello.db")
	app.MbinServe(":3033")
}
