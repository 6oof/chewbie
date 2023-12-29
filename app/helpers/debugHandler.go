package helpers

import (
	"fmt"
	"net/http"

	"github.com/davecgh/go-spew/spew"
)

// DDH is a debugging helper function that pretty-prints the variable and returns a simple HTML response.
func DDH(w http.ResponseWriter, v interface{}) {
	// Convert the variable to a formatted HTML string
	htmlString := spew.Sdump(v)

	// Create a simple HTML response
	responseHTML := fmt.Sprintf(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>DDH Debug</title>
		</head>
					 <body style="background-color: #30363d; color: white; padding: 10px;">
			<h1>DDH Debug Output</h1>
			<pre>%s</pre>
		</body>
		</html>
	`, htmlString)

	// Write the HTML response to the http.ResponseWriter
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(responseHTML))
	panic("Exectution stopped by DDH")
}
