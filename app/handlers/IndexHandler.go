package handlers

import (
	"net/http"

	minitemp "github.com/6oof/miniweb-base/app/templateEngine"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	t := minitemp.PageTemplate{
		Files: []string{"layout", "index/index"},
		Seo:   minitemp.Seo{Title: "Home Page"},
		Data:  nil,
	}

	t.RenderPageAndSend(w)
}
