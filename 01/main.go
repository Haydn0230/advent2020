package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main () {
	dataBytes, err := ioutil.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}

	data := make([]int,0)
	json.Unmarshal(dataBytes, &data)

	fmt.Println(data)
}
