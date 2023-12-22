package handlers

import (
	"net/http"

	minitemp "github.com/6oof/miniweb-base/app/templateEngine"
)

func HandleNotFound(w http.ResponseWriter, r *http.Request) {
	t := minitemp.PageTemplate{
		Layout: "empty",
		Page:   "error",
		Seo: minitemp.Seo{
			Title: "Page Not Found",
		},
		Data: map[string]interface{}{
			"code":    404,
			"message": "Sorry, we can't find that page. You'll find lots to explore on the home page.",
		},
	}

	t.RenderPageAndSend(w)
}
