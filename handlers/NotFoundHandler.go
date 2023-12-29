package handlers

import (
	"net/http"

	"github.com/6oof/miniweb-base/app/templateEngine"
)

func HandleNotFound(w http.ResponseWriter, r *http.Request) {
	t := mwtemp.PageTemplate{
		Layout: "empty",
		Page:   "error",
		Seo: mwtemp.Seo{
			Title: "Page Not Found",
		},
		Data: map[string]interface{}{
			"code":    404,
			"message": "Sorry, we can't find that page. You'll find lots to explore on the home page.",
		},
	}

	t.RenderPageAndSend(w, r)
}
