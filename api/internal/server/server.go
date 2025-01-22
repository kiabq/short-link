package server

import (
	"kiabq/hyperslice/internal/routes"
	"log"
	"net/http"
	"time"
)

func Start() {
	s := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	routes.RegisterRoutes()

	log.Fatal(s.ListenAndServe())
}
