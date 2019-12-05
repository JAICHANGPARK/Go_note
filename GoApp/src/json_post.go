package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Persion struct {
	Name string
	Age  int
}

func main() {
	persion := Persion{"Alex", 10}
	pbytes, _ := json.Marshal(persion)
	buff := bytes.NewBuffer(pbytes)
	resp, err := http.Post("http://httpbin.org/post", "application/json", buff)
}
