package main

import (
	"log"
	"net/http"
	"friday/api"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("/var/www/gofridayadmin/")))
	http.HandleFunc("/api/", api.HandleApiSlash)
	http.HandleFunc("/api/get/", api.HandleApiGet)
	http.HandleFunc("/api/set/", api.HandleApiSet)
	log.Fatal(http.ListenAndServe("0.0.0.0:7008", nil))
}
