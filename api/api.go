package api

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

var ram []Memory = make([]Memory, 0)

func search(username string, application string, key string, ram []Memory) []Memory {
	var result []Memory
	for _, item := range ram {
		if item.Owner.username == username && item.App.application == application && item.Object.key == key {
			result = append(result, item)
		}
	}
	return result
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

	if searchres == nil {
		fmt.Println("INVALID GET FOR " + username + ":" + application + "{" + key + "}")
		fmt.Fprintln(w, "Not found!")
	}

	for _, item := range searchres {
		fmt.Println("GET for " + item.username + ":" + item.application + "; {" + item.page + ":'" + item.key + "'->'" + item.value + "'} at " + strconv.FormatInt(time.Now().UnixNano(), 10))
		fmt.Fprintln(w, "GET for "+item.username+":"+item.application+"; {"+item.page+":'"+item.key+"'->'"+item.value+"'}<br>")
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

	fmt.Fprintln(w, "SET for "+username+":"+application+"; {"+page+":'"+key+"'->'"+value+"'} at "+strconv.FormatInt(item.Object.time, 10))
	fmt.Println("SET for " + username + ":" + application + "; {" + page + ":'" + key + "'->'" + value + "'} at " + strconv.FormatInt(item.Object.time, 10))
}
