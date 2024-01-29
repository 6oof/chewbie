package handlers

import (
	"context"
	"net/http"

	"github.com/6oof/minweb/app/helpers"
	"github.com/6oof/minweb/app/views/pages"
)

func HandleNotFound(w http.ResponseWriter, r *http.Request) {
	seo := helpers.BaseSeo()
	seo.Title = "404"
	pg := pages.ErrorPage(seo, "404", "Sorry, the page you are looking for does not exist.")
	pg.Render(context.Background(), w)
}
