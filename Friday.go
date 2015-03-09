package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Message struct {
	Key   string
	Value string
	Time  int64
}

func handle_slash(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/static/", 301)
}
func handle_api_slash(w http.ResponseWriter, r *http.Request) {
	m := Message{"_error", "Incomplete API request.", 0}
	b, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "%s", b)
}
func handle_api_get(w http.ResponseWriter, r *http.Request) {
}
func handle_api_set(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // Parses the request body
	key := r.PostFormValue("Key")
	value := r.PostFormValue("Value")

	m := Message{key, value, time.Now().UnixNano()}
	b, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "%s", b)
	// html.EscapeString()
}
func main() {
	http.HandleFunc("/", handle_slash)
	http.HandleFunc("/api/", handle_api_slash)
	http.HandleFunc("/api/get/", handle_api_get)
	http.HandleFunc("/api/set/", handle_api_set)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Fatal(http.ListenAndServe("0.0.0.0:7008", nil))
}
