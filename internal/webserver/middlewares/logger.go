package middlewares

import (
	"fmt"
	"log"
	"net/http"
)

func Logging(next http.Handler) http.Handler {
	fmt.Println("hereeeee")
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Do stuff here
        log.Println(r.RequestURI)
        // Call the next handler, which can be another middleware in the chain, or the final handler.
        next.ServeHTTP(w, r)
    })
}