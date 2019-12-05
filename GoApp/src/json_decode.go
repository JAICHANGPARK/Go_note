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
	jsonBytes, _ := json.Marshal(Member{"Tim", 1, true})
	println(jsonBytes)

	var mem Member
	err := json.Unmarshal(jsonBytes, &mem)

	if err != nil {
		panic(err)
	}

	fmt.Println(mem.Name, mem.Age, mem.Active)

}
