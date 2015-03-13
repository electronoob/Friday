package api

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
