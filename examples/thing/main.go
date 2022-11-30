package main

import (
	"encoding/json"
	"fmt"

	"github.com/hiroaqii/go-bgg/bgg"
)

func main() {
	hotItems, err := bgg.Thing()
	if err != nil {
		println(err)
		return
	}

	e, _ := json.Marshal(hotItems)
	fmt.Println(string(e))

}
