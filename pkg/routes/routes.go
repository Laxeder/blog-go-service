package routes

import (
	"net/http"
)

func Routes() *http.ServeMux {
	mux := &http.ServeMux{}

	mux.HandleFunc("/hello", Hello)
	mux.HandleFunc("/", NotFound)

	return mux
}
