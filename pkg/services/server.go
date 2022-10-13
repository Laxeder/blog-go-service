package services

import (
	"blog/pkg/routes"
	"net/http"
)

func Server() http.Server {
	routes := routes.Routes()

	srv := http.Server{
		Handler: routes,
	}

	return srv
}
