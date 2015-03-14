package main

import (
	"friday/api"
	"friday/backup"
	"friday/ui"
	"log"
	"net/http"
)

func main() {
	r := ui.BoolPrompt("Do you want to restore from backup?")
	if r {
		api.Ram = backup.ReadDump()
	}
	b := ui.BoolPrompt("Do you want to keep backups?")
	if b {
		t := ui.IntPrompt("How often? (in minutes)")
		go backup.SDump(t)
	}
	http.Handle("/", http.FileServer(http.Dir("/var/www/gofridayadmin/")))
	http.HandleFunc("/api/", api.HandleApiSlash)
	http.HandleFunc("/api/get/", api.HandleApiGet)
	http.HandleFunc("/api/set/", api.HandleApiSet)
	log.Fatal(http.ListenAndServe("0.0.0.0:7008", nil))
}
