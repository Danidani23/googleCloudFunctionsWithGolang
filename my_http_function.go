package mycloudeventfunction

import (
	"fmt"
	"net/http"
)

// helloWorld writes "Hello, World!" to the HTTP response.
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintln(w, "Hello, World!")
}
