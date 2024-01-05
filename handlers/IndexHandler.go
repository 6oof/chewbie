package handlers

import (
	"context"
	"net/http"

	"github.com/6oof/miniweb-base/app/helpers"
	"github.com/6oof/miniweb-base/views/pages"
	"github.com/gorilla/csrf"
)

// HandleIndex is the handler function for the "/" route, rendering the home page.
func HandleIndex(w http.ResponseWriter, r *http.Request) {
	// Define the page template for the home page
	seo := helpers.BaseSeo()
	pg := pages.Index(seo, csrf.TemplateField(r), "", "")
	pg.Render(context.Background(), w)
}
