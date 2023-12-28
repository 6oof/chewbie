package mwtemp

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"

	"github.com/6oof/miniweb-base/app/helpers"
)

type Seo struct {
	Name        string // if none is given it will be taken from .env NAME field
	Title       string
	Description string // if none is given it will be taken from .env DESCRIPTION field
	Keywords    string
}

// PageTemplate represents a structured template for rendering entire pages. It encapsulates the necessary information to render a page,
// including the list of template files and additional data to be used in the template execution.
//
// Example Usage:
//
//	minitemp.PageTemplate{
//		Layout:     "layout",
//		Page:       "index",
//		Components: []string{"navbar", "footer"},
//		Seo:        minitemp.Seo{Title: "Home Page"},
//		Data:       nil,
//	}
type PageTemplate struct {
	Layout     string      // Optional layout file to be rendered form ./views/layouts without file extension;
	Page       string      //File name of the page to be rendered inside ./views/pages without file extension;
	Components []string    // Relative paths to template files of components inside ./views/components without file extension;
	Seo        Seo         // SEO data
	Data       interface{} // Data to be used in the template
}

// FragmentTemplate represents a template for rendering fragments of a page. It specifies the list of template files, the outermost block
// to be rendered, and any additional data required for template execution.
//
// Example Usage:
//
//	minitemp.FragmentTemplate{
//		Files:     []string{"components/navbar"}
//		BlockName: "navbar",
//		Data:      nil,
//	}
type FragmentTemplate struct {
	Files     []string    // Relative paths to template files (excluding trailing slash and extension);
	BlockName string      // The name of the outermost defined block to be rendered
	Data      interface{} // Data to be used in the template
}

// RenderPage generates the HTML content for the specified PageTemplate and returns it as a string. Any data can be accesed in the template via {{.Data}}
func (p PageTemplate) RenderPage() (string, error) {
	// Prepend the correct path to the template files
	for i, v := range p.Components {
		p.Components[i] = "./views/components/" + v + ".go.html"
	}

	// Set default layout if not provided
	if p.Layout == "" {
		p.Layout = "./views/layouts/layout.go.html"
	} else {
		p.Layout = "./views/layouts/" + p.Layout + ".go.html"
	}

	p.Page = "./views/pages/" + p.Page + ".go.html"

	// Include the superglobals file
	sg := "./views/superglobals.go.html"
	completeFiles := append([]string{sg, p.Layout, p.Page}, p.Components...)

	// Create a new template instance
	ts := template.New("")

	// Register custom template functions
	ts.Funcs(CustomTemplateFuncs)

	// Parse the template files
	ts, err := ts.ParseFiles(completeFiles...)
	if err != nil {
		return "", fmt.Errorf("error parsing template files: %v", err)
	}

	// Set default SEO values if not provided
	if p.Seo.Name == "" {
		p.Seo.Name = helpers.EnvOrPanic("NAME")
	}
	fmt.Print(p.Seo.Name)

	if p.Seo.Description == "" {
		p.Seo.Description = helpers.Env("DESCRIPTION", "App created with Miniweb")
	}

	data := map[string]interface{}{
		"Seo":  p.Seo,
		"Data": p.Data,
	}

	var contentBuffer bytes.Buffer
	err = ts.ExecuteTemplate(&contentBuffer, "layout", data)

	return contentBuffer.String(), nil
}

// RenderPageAndSend generates the HTML content for the specified PageTemplate and sends it as an HTTP response.
func (p PageTemplate) RenderPageAndSend(w http.ResponseWriter) {
	html, err := p.RenderPage()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(fmt.Sprint(err)))
	}
	// Send the HTML response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html))
}

// RenderFragment generates the HTML content for the specified FragmentTemplate and returns it as a string.
func (p FragmentTemplate) RenderFragment() (string, error) {
	// Prepend the correct path to the template files
	for i, v := range p.Files {
		p.Files[i] = "./views/" + v + ".go.html"
	}

	// Create a new template instance
	ts := template.New("")

	// Register custom template functions
	ts.Funcs(CustomTemplateFuncs)

	// Parse the template files
	ts, err := ts.ParseFiles(p.Files...)
	if err != nil {
		return "", fmt.Errorf("error parsing template files: %v", err)
	}

	var contentBuffer bytes.Buffer
	err = ts.ExecuteTemplate(&contentBuffer, p.BlockName, p.Data)

	return contentBuffer.String(), nil
}

// RenderFragmentAndSend generates the HTML content for the specified FragmentTemplate and sends it as an HTTP response.
func (p FragmentTemplate) RenderFragmentAndSend(w http.ResponseWriter) {
	html, err := p.RenderFragment()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(fmt.Sprint(err)))
	}
	// Send the HTML response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html))
}
