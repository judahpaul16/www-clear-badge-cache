package main

import (
	"log"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		// Log the incoming request
		log.Printf("Received %s request for %s from %s", r.Method, r.URL.Path, r.RemoteAddr)

		next.ServeHTTP(w, r) // pass control to the next handler

		// Log the completion of the handling
		log.Printf("Completed %s request for %s in %v", r.Method, r.URL.Path, time.Since(start))
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", serveTemplate)
	mux.HandleFunc("/clear-cache", clearCache)

	// Wrap the ServeMux with the LoggingMiddleware
	loggedMux := LoggingMiddleware(mux)

	log.Fatal(http.ListenAndServe(":8080", loggedMux))
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	home := "index.html"
	http.ServeFile(w, r, home)
}

func clearCache(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	log.Println("Cache cleared for URL:", url)
	w.Write([]byte("Cache cleared for " + url))
}
