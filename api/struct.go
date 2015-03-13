package api

import (
	"github.com/pmylund/go-cache"
	"time"
)

type Message struct {
	Key   string
	Value string
	Time  int64
}

var c = cache.New(5*time.Hour, 1*time.Hour)
var i int = 0
