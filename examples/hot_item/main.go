package main

import (
	"encoding/json"
	"fmt"

	"github.com/hiroaqii/bgggo"
)

func main() {
	hotItems, err := bgggo.Hot()
	if err != nil {
		println(err)
		return
	}

	e, _ := json.Marshal(hotItems)
	fmt.Println(string(e))

}
