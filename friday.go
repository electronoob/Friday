package main

import (
	"flag"
	"friday/api"
	"friday/backup"
	"log"
	"net/http"
)

func main() {
	dir := flag.String("d", "/var/www/gofridayadmin/", "Directory to serve from")
	addr := flag.String("l", "0.0.0.0:7008", "Listening address")
	file := flag.String("f", "friday.dump", "File to restore backup from and save to")
	r := flag.Bool("r", false, "Bool: Restore from backup")
	t := flag.Int("b", 1, "Backup frequency (in minutes, or 0 to disable)")

	flag.Parse()

	if *r {
		api.Ram = backup.ReadDump(*file)
	}
	if *t != 0 {
		go backup.SDump(*t, *file)
	}

	http.Handle("/", http.FileServer(http.Dir(*dir)))
	http.HandleFunc("/api/", api.HandleApiSlash)
	http.HandleFunc("/api/get/", api.HandleApiGet)
	http.HandleFunc("/api/set/", api.HandleApiSet)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
