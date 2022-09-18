package bgggo

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	xj "github.com/basgys/goxml2json"
)

type HotItems struct {
	Items struct {
		Termsofuse string `json:"-termsofuse"`
		Item       []struct {
			Rank      string `json:"-rank"`
			Thumbnail struct {
				Value string `json:"-value"`
			} `json:"thumbnail"`
			Name struct {
				Value string `json:"-value"`
			} `json:"name"`
			Yearpublished struct {
				Value string `json:"-value"`
			} `json:"yearpublished"`
			ID string `json:"-id"`
		} `json:"item"`
	} `json:"items"`
}

func Hot() (HotItems, error) {
	resp, err := http.Get("https://api.geekdo.com/xmlapi2/hot")
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	sb := string(body)
	xml := strings.NewReader(sb)
	jsonStr, err := xj.Convert(xml)
	if err != nil {
		panic("That's embarrassing...")
	}

	var hotItems HotItems
	if err := json.Unmarshal([]byte(jsonStr.Bytes()), &hotItems); err != nil {
		fmt.Println(err)
		return HotItems{}, err
	}

	return hotItems, nil
}
