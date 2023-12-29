package templatefuncs

import "github.com/6oof/miniweb-base/app/helpers"

func ProjectName() string {
	return helpers.Env("NAME", "MiniWeb")
}
