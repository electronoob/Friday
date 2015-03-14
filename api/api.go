package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

var ram []Memory = make([]Memory, 0)

func search(username string, application string, key string, ram []Memory) []Memory {
	var result []Memory
	for _, item := range ram {
		if item.Owner.Username == username && item.App.Application == application && item.Object.Key == key {
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
	} else {

		j, err := json.Marshal(searchres)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(j)

		for _, item := range searchres {
			fmt.Println("GET for " + item.Username + ":" + item.Application + "; {" + item.Page + ":'" + item.Key + "'->'" + item.Value + "'} at " + strconv.FormatInt(time.Now().UnixNano(), 10))
		}
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

	item.Owner.Username = username
	item.App.Application = application
	item.Object.Page = page
	item.Object.Key = key
	item.Object.Value = value
	item.Object.Time = time.Now().UnixNano()

	ram = append(ram, item)

	fmt.Fprintln(w, "SET for "+username+":"+application+"; {"+page+":'"+key+"'->'"+value+"'} at "+strconv.FormatInt(item.Object.Time, 10))
	fmt.Println("SET for " + username + ":" + application + "; {" + page + ":'" + key + "'->'" + value + "'} at " + strconv.FormatInt(item.Object.Time, 10))
}
