package templatefuncs

import "time"

func CurrentYear() int {
	currentYera := time.Now().Year()
	return currentYera
}
