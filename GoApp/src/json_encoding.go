package main

import (
	"encoding/json"
	"fmt"
)

//Member -
type Member struct {
	Name   string
	Age    int
	Active bool
}

func main() {
	mem := Member{"Alex", 10, true}
	jsonBytes, err := json.Marshal(mem)
	if err != nil {
		panic(err)
	}
	jsonString := string(jsonBytes)
	fmt.Println(jsonString)
}
