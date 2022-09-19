package middlewares

import (
	"log"
	"net/http"
)

func Logging(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("received %s request : %s", r.Method, r.RequestURI)
        // Call the next handler, which can be another middleware in the chain, or the final handler.
        next.ServeHTTP(w, r)
    })
}