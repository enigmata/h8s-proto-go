package v1

import (
	"net/http"
)

func InstallHandlers(mux *http.ServeMux) {
	mux.HandleFunc(UriPathAdminVersion, AdminVersionHandler)
	return
}
