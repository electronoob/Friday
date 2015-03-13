package main

import (
	"fmt"
	"time"
)

func main() {
	username := "electronoob"
	application := "DEADBEEF"
	page := "AEEEEE"
	key := "roar"
	value := "this sure is complicated!"

	type Owner struct{ username string }
	type App struct{ application string }
	type Object struct {
		key   string
		value string
		time  int64
		page  string
	}
	type Memory struct {
		Owner
		App
		Object
	}
	ram := make([]Memory, 0)

	item := Memory{}
	item.Owner.username = username
	item.App.application = application
	item.Object.page = page
	item.Object.key = key
	item.Object.value = value
	item.Object.time = time.Now().UnixNano()

	ram = append(ram, item)
	fmt.Printf("%v\n", ram)
}
