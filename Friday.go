package main

import (
	"encoding/json"
	"fmt"
	"github.com/pmylund/go-cache"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Message struct {
	Key   string
	Value string
	Time  int64
}

var c = cache.New(5*time.Hour, 1*time.Hour)
var i int = 0

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
	r.ParseForm()
	result := r.URL.Path[9:len(r.URL.Path)]
	gresult, found := c.Get(result)
	if !found {
		fmt.Fprintln(w, "Not found!")
	} else {
		fmt.Fprintf(w, "%s", gresult)
	}

}
func handle_api_set(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // Parses the request body
	key := r.PostFormValue("Key")
	value := r.PostFormValue("Value")

	m := Message{key, value, time.Now().UnixNano()}
	json, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}
	inum := strconv.Itoa(i)
	c.Set(inum, json, cache.DefaultExpiration)
	rvalue := ("entry " + inum + " with key-value pair <br>'" + m.Key + "'->'" + m.Value + "'<br>at time " + strconv.FormatInt(m.Time, 10))
	fmt.Fprintln(w, rvalue)
	i++
}
func main() {
	http.HandleFunc("/", handle_slash)
	http.HandleFunc("/api/", handle_api_slash)
	http.HandleFunc("/api/get/", handle_api_get)
	http.HandleFunc("/api/set/", handle_api_set)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Fatal(http.ListenAndServe("0.0.0.0:7008", nil))
}
