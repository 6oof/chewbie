package minitemp

import (
	"html/template"

	templatefuncs "github.com/6oof/miniweb-base/app/templateEngine/funcs"
)

var CustomTemplateFuncs = template.FuncMap{
	"customUppercase": templatefuncs.CustomUppercase,
}
