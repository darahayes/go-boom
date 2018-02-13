package boom

import (
	"log"
	"net/http"
	"runtime/debug"
)

// RecoverHandler is a middleware handler function that can be used
// to recover from unexpected panics, log a stack trace and respond with a generic
// 500 Internal Server Error.
// Ensures no sensitive data is leaked during panics.
func RecoverHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %+v", err)
				debug.PrintStack()
				BadImplementation(w)
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

// NotFoundHandler handler function that uses boom.NotFound() to create a structured 404 response.
// Can be used to configure routers that accept a handler for 404 cases
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	NotFound(w)
}
