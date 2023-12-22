package handlers

import (
	"net/http"

	minitemp "github.com/6oof/miniweb-base/app/templateEngine"
)

// HandleIndex is the handler function for the "/" route, rendering the home page.
func HandleIndex(w http.ResponseWriter, r *http.Request) {
	// Define the page template for the home page
	t := minitemp.PageTemplate{
		Page: "index",
		Seo:  minitemp.Seo{Title: "Home Page"},
		Data: nil,
	}

	// Render the page and send it as an HTTP response
	t.RenderPageAndSend(w)
}
