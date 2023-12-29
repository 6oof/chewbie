package handlers

import (
	"net/http"

	mwtemp "github.com/6oof/miniweb-base/app/templateEngine"
)

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
