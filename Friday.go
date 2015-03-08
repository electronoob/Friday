package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "log"
)
type Message struct {
    Key string
    Value string
    Time int64
}

func handle_slash(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, "/static/", 301)
}

func handle_db(w http.ResponseWriter, r *http.Request) {
 m := Message{"Alice", "Hello", 1294706395881547000}
 b, err := json.Marshal(m)
 if err != nil {
    log.Fatal(err)
 }
 fmt.Fprintf(w, "%s", b)
}
func main() {
    http.HandleFunc("/", handle_slash)
    http.HandleFunc("/db/", handle_db)
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
    http.ListenAndServe("0.0.0.0:7008", nil)
}
