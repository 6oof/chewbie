package handlers

import (
	"net/http"

	"github.com/6oof/minweb/app/helpers"
	"github.com/6oof/minweb/app/helpers/temp"
)

type NotFoundData struct {
	Code    int
	Message string
}

func HandleNotFound(w http.ResponseWriter, r *http.Request) {
	seo := helpers.BaseSeo()
	seo.Title = "Page Not Found"

	t := temp.PageTemplate{
		Layout: "empty",
		Page:   "error",
		Seo:    seo,
		Data: NotFoundData{
			Code:    404,
			Message: "The page you are looking for does not exist.",
		},
	}

	t.RenderPageAndSend(w, r)
}
