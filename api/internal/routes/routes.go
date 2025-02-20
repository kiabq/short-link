package routes

import (
	"errors"
	"fmt"
	"html"
	"net/http"
)

const (
	INVALID_METHOD string = "Method Not Allowed"
)

func RegisterRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Cannot use any other method(s) to access this route
		if r.Method != "GET" {
			err := errors.New(INVALID_METHOD)
			http.Error(w, err.Error(), 405)
			return
		}
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" || r.Method != "GET" {
			err := errors.New(INVALID_METHOD)
			http.Error(w, err.Error(), 405)
			return
		}

		if r.Method == "GET" {
			// Do GET logic
		}

		if r.Method == "POST" {
			// Do POST logic
		}

		id := r.PathValue("id")
		fmt.Fprintf(w, "Hello, path: %q, id: %s", html.EscapeString(r.URL.Path), id)
	})
}
