package main

import (
	"encoding/json"
	"fmt"

	"github.com/hiroaqii/bgggo"
)

func main() {
	collectioItems, err := bgggo.Collection()
	if err != nil {
		println(err)
		return
	}

	e, _ := json.Marshal(collectioItems)
	fmt.Println(string(e))

	//for _, item := range collectioItems.Items {
	//	fmt.Println(item)
	//}
}
