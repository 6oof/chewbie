package minitemp

import (
	"html/template"

	templatefuncs "github.com/6oof/miniweb-base/app/templateEngine/funcs"
)

// CustomTemplateFuncs is a FuncMap that holds custom template functions used in rendering templates.
var CustomTemplateFuncs = template.FuncMap{
	// An example - implementation can be found in app/templateEngine/funcs/customUppercase.go
	"CurrentYear": templatefuncs.CurrentYear,
}
