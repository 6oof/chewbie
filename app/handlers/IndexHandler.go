package handlers

import (
	"context"
	"net/http"

	"github.com/6oof/minweb/app/helpers"
	"github.com/6oof/minweb/app/views/components"
	"github.com/6oof/minweb/app/views/pages"
	"github.com/gorilla/csrf"
)

// HandleIndex is the handler function for the "/" route, rendering the home page.
func HandleIndex(w http.ResponseWriter, r *http.Request) {
	// Define the page template for the home page
	seo := helpers.BaseSeo()
	pg := pages.Index(seo, csrf.TemplateField(r), "", "")
	pg.Render(context.Background(), w)
}

type ShowcaseFormData struct {
	Result    string
	NameError string
}

func HandleShowcaseFormPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Sorry, something went wrong", http.StatusInternalServerError)
	}

	res := ShowcaseFormData{}

	if len(r.PostForm.Get("Name")) < 1 {
		res.NameError = "Name is required"
	} else {
		if r.PostForm.Get("Name") == "Bruce Wayne" {
			res.Result = "You're Batman"
		} else {
			res.Result = "You're not Batman"
		}
	}
	frag := components.ShowcaseForm(csrf.TemplateField(r), res.NameError, res.Result)
	frag.Render(context.Background(), w)

}
