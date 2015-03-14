package backup

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"friday/api"
	"io/ioutil"
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

func writeDump(m []api.Memory) {
	b, _ := getBytes(m)
	err := ioutil.WriteFile("friday.dump", b, 0664)
	if err != nil {
		panic(err)
	}
}

func ReadDump() []api.Memory {
	f, err := ioutil.ReadFile("friday.dump")
	if err != nil {
		panic(err)
	}
	dec := gob.NewDecoder(bytes.NewReader(f))
	var ret []api.Memory
	err = dec.Decode(&ret)
	if err != nil {
		panic(err)
	}
	return ret
}

func SDump(t int) {
	i := 0
	for i == 0 {
		time.Sleep(time.Minute * time.Duration(t))
		writeDump(api.Ram)
		fmt.Println("BACKUP")
	}
}
