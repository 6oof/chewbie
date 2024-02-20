package templatefuncs

import "github.com/6oof/minweb/app/helpers"

func ProjectName() string {
	return helpers.Env("NAME", "MiniWeb")
}
