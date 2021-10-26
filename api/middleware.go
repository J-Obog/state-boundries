package api

import (
	"log"
	"net/http"
)

// middleware for logging requests
func ReqLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

// CORS middleware
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Origin, Content-Type, Authorization, cache-control")
		w.Header().Set("Access-Control-Allow-Credentials", "false")
		w.Header().Set("Content-Type", "application/json")

		// handle OPTIONS request
		if r.Method == "OPTIONS" {
			w.WriteHeader(200)
			return
		}
		
		next.ServeHTTP(w, r)
	})
}