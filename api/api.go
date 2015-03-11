package api

import (
	"encoding/json"
	"fmt"
	"github.com/pmylund/go-cache"
	"log"
	"net/http"
	"strconv"
	"time"
)

func HandleApiSlash(w http.ResponseWriter, r *http.Request) {
	m := Message{"_error", "Incomplete API request.", 0}
	b, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "%s", b)
}
func HandleApiGet(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	result := r.URL.Path[9:len(r.URL.Path)]
	gresult, found := c.Get(result)
	if !found {
		fmt.Fprintln(w, "Not found!")
		fmt.Println("INVALUD GET for " + result)
	} else {
		fmt.Fprintf(w, "%s", gresult)
		fmt.Println("GET for " + result)
	}
}
func HandleApiSet(w http.ResponseWriter, r *http.Request) {
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
	fmt.Println("SET #" + inum + "; " + m.Key + "'->'" + m.Value + "'; " + strconv.FormatInt(m.Time, 10))
	i++
}
