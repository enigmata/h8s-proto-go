package v1

import (
	"fmt"
	"net/http"
)

func AdminVersionHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "version 1.0")
	return
}
