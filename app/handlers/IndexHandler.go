package handlers

import (
	"net/http"

	"github.com/6oof/minweb/app/helpers"
	"github.com/6oof/minweb/app/helpers/temp"
)

// HandleIndex is the handler function for the "/" route, rendering the home page.
func HandleIndex(w http.ResponseWriter, r *http.Request) {
	// Define the page template for the home page
	seo := helpers.BaseSeo()
	seo.Title = "Home Page"

	t := temp.PageTemplate{
		Page:       "index",
		Components: []string{"showcaseForm"},
		Seo:        seo,
	}

	// Render the page and send it as an HTTP response
	t.RenderPageAndSend(w, r)
}

type ShowcaseFormResult struct {
	NameError string
	Result    string
}

func HandleShowcaseFormPost(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()

	if err != nil {
		http.Error(w, "Sorry, something went wrong", http.StatusInternalServerError)
	}
	data := ShowcaseFormResult{}

	if len(r.PostForm.Get("Name")) < 1 {
		data.NameError = "Name is required"
	} else {
		if r.PostForm.Get("Name") == "Bruce Wayne" {
			data.Result = "You're Batman"
		} else {
			data.Result = "You're not Batman"
		}
	}

	t := temp.FragmentTemplate{
		Files:     []string{"components/showcaseForm"},
		BlockName: "showcaseForm",
		Data:      data,
	}

	t.RenderFragmentAndSend(w, r)

}
