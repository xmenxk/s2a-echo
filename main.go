// Package main runs an AppEngine app that queries the Bifrost Echo service.
package main

import (
	"google.golang.org/appengine"
	"log"
	"net/http"
	"os"
)

func main() {
	// flag.Set("s2a_enable_appengine_dialer", "true")
	http.HandleFunc("/", DoEcho)
	http.HandleFunc("/v2", DoEchoV2)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, appengine.Middleware(http.DefaultServeMux)); err != nil {
		log.Fatal(err)
	}
}
