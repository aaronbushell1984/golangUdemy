// Package context demonstrates use of context to propagate cancellation of go routines
package context

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

// Store defines behaviour to fetch and cancel
//
// Context is passed as an argument which is idiomatic and mandated by google for all its programs
type Store interface {
	Fetch(ctx context.Context) (string, error)
	Cancel()
}

// Server accepts a store to pass to and return a function which uses a writer to print a requests' response
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())
		if err != nil {
			log.Println("error fetching from store")
		}
		_, err = fmt.Fprint(w, data)
		if err != nil {
			log.Println("error writing data")
		}
	}
}
