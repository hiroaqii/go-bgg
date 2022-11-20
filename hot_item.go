package bgggo

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

type hotItemsIntermediate struct {
	XMLName    xml.Name `xml:"items"`
	Text       string   `xml:",chardata"`
	Termsofuse string   `xml:"termsofuse,attr"`
	Item       []struct {
		Text      string `xml:",chardata"`
		ID        string `xml:"id,attr"`
		Rank      string `xml:"rank,attr"`
		Thumbnail struct {
			Text  string `xml:",chardata"`
			Value string `xml:"value,attr"`
		} `xml:"thumbnail"`
		Name struct {
			Text  string `xml:",chardata"`
			Value string `xml:"value,attr"`
		} `xml:"name"`
		Yearpublished struct {
			Text  string `xml:",chardata"`
			Value string `xml:"value,attr"`
		} `xml:"yearpublished"`
	} `xml:"item"`
}

type HotItems struct {
	Items []HotItem `json:"items"`
}

type HotItem struct {
	ID            int    `json:"id"`
	Rank          int    `json:"rank"`
	Name          string `json:"name"`
	Yearpublished int    `json:"yearpublished,omitempty"`
	Thumbnail     string `json:"thumbnail"`
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

	var itemInter hotItemsIntermediate
	err = xml.Unmarshal(body, &itemInter)
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	var hotItems = []HotItem{}
	for _, item := range itemInter.Item {
		id, _ := strconv.Atoi(item.ID)
		rank, _ := strconv.Atoi(item.Rank)
		yearpublished, _ := strconv.Atoi(item.Yearpublished.Value)

		hotItems = append(hotItems, HotItem{
			ID:            id,
			Rank:          rank,
			Thumbnail:     item.Thumbnail.Value,
			Name:          item.Name.Value,
			Yearpublished: yearpublished,
		})
	}

	return HotItems{Items: hotItems}, nil
}
