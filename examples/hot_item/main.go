package main

import (
	"fmt"
	"strconv"

	"github.com/hiroaqii/bgggo"
)

func main() {
	hotItems, err := bgggo.Hot()
	if err != nil {
		println(err)
		return
	}

	for _, item := range hotItems.Items.Item {
		rank, _ := strconv.Atoi(item.Rank)
		fmt.Printf("%02d,%s(%s)\n", rank, item.Name.Value, item.Yearpublished.Value)
	}
	for _, item := range hotItems.Items.Item {
		fmt.Println(item)
	}
}
