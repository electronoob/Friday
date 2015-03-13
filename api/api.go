package api

import (
	"fmt"
	"net/http"
	"time"
)

var ram []Memory = make([]Memory, 0)

func search(username string, application string, key string, ram []Memory) Memory {
	for _, item := range ram {
		if item.Owner.username == username && item.App.application == application && item.Object.key == key {
			return item
		}
	}
	return Memory{}
}

func HandleApiSlash(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "INVALID GET for /api/")
	fmt.Println("INVALID GET for /api/")
}
func HandleApiGet(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // Parses the request body

	username := r.PostFormValue("User")
	application := r.PostFormValue("App")
	key := r.PostFormValue("Key")

	searchres := search(username, application, key, ram)
	emptyMem := Memory{}
	if searchres != emptyMem {
		fmt.Fprintf(w, "GET for " + searchres.username + ":" + searchres.application + "; {" + searchres.page + ":'" + searchres.key + "'->'" + searchres.value + "'}")
	} else {
		fmt.Fprintf(w, "Not found!")
	}
}
func HandleApiSet(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // Parses the request body

	username := r.PostFormValue("User")
	application := r.PostFormValue("App")
	page := r.PostFormValue("Page")
	key := r.PostFormValue("Key")
	value := r.PostFormValue("Value")

	item := Memory{}
	item.Owner.username = username
	item.App.application = application
	item.Object.page = page
	item.Object.key = key
	item.Object.value = value
	item.Object.time = time.Now().UnixNano()

	ram = append(ram, item)
	fmt.Fprintf(w, "SET for " + username + ":" + application + "; {" + page + ":'" + key + "'->'" + value + "'}")
}
