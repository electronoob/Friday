package api

type Owner struct{ Username string }
type App struct{ Application string }
type Object struct {
	Key   string
	Value string
	Time  int64
	Page  string
}
type Memory struct {
	Owner
	App
	Object
}
