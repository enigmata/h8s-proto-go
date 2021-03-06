package v1

import (
	"fmt"
	"log"
	"net/http"

	"github.com/enigmata/h8s-proto-go/version"
)

func AdminVersionHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("INFO: %s %s", r.Method, r.RequestURI)
	fmt.Fprintf(w, version.GetVersion().Marshal())
	return
}
