package middleware

import (
	"log"
	"net/http"
)

func Loader(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n%s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}
