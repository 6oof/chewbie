package handlers

import (
	"net/http"

	mwtemp "github.com/6oof/minweb/app/helpers/templateEngine"
)

// HandleIndex is the handler function for the "/" route, rendering the home page.
func HandleIndex(w http.ResponseWriter, r *http.Request) {
	// Define the page template for the home page
	t := mwtemp.PageTemplate{
		Page:       "index",
		Seo:        mwtemp.Seo{Title: "Home Page"},
		Components: []string{"showcaseForm"},
		Data:       nil,
	}

	// Render the page and send it as an HTTP response
	t.RenderPageAndSend(w, r)
}

type ShowcaseForm struct {
	Name string
}

type Response struct {
	Result string
	Errors map[string]string
}

func HandleShowcaseFormPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Sorry, something went wrong", http.StatusInternalServerError)
	}

	res := Response{
		Errors: map[string]string{},
	}

	if len(r.PostForm.Get("Name")) < 1 {
		res.Errors["Name"] = "Name is required"
	} else {
		if r.PostForm.Get("Name") == "Bruce Wayne" {
			res.Result = "You're Batman"
		} else {
			res.Result = "You're not Batman"
		}
	}

	t := mwtemp.FragmentTemplate{
		Files:     []string{"components/showcaseForm"},
		BlockName: "showcaseForm",
		Data:      res,
	}

	t.RenderFragmentAndSend(w, r)

}
