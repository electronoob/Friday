package backup

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"friday/api"
	"io/ioutil"
	"strconv"
	"time"
)

func getBytes(key interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func writeDump(m []api.Memory, filename string) {
	b, _ := getBytes(m)
	err := ioutil.WriteFile(filename, b, 0664)
	if err != nil {
		panic(err)
	}
}

func ReadDump(filename string) []api.Memory {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("File", filename, "not found!")
		return []api.Memory{}
	}
	dec := gob.NewDecoder(bytes.NewReader(f))
	var ret []api.Memory
	err = dec.Decode(&ret)
	if err != nil {
		fmt.Println(err)
	}
	return ret
}

func SDump(t int, filename string) {
	for {
		time.Sleep(time.Minute * time.Duration(t))
		writeDump(api.Ram, filename)
		fmt.Println("BACKUP at " + strconv.FormatInt(time.Now().UnixNano(), 10))
	}
}
