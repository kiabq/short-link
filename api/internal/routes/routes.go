package routes

import (
	"fmt"
	"html"
	"net/http"
)

func RegisterRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			// Return error
		}
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			// Return error
		}

		id := r.PathValue("id")
		fmt.Fprintf(w, "Hello, path: %q, id: %s", html.EscapeString(r.URL.Path), id)
	})
}
