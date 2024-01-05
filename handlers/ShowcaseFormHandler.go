package handlers

import (
	"context"
	"net/http"

	"github.com/6oof/miniweb-base/views/components"
	"github.com/gorilla/csrf"
)

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
