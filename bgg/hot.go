package bgg

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

type hotItemsIntermediate struct {
	XMLName    xml.Name              `xml:"items"`
	Text       string                `xml:",chardata"`
	Termsofuse string                `xml:"termsofuse,attr"`
	Item       []hotItemIntermediate `xml:"item"`
}

type hotItemIntermediate struct {
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

func newHotItem(xml hotItemIntermediate) HotItem {

	id, _ := strconv.Atoi(xml.ID)
	rank, _ := strconv.Atoi(xml.Rank)
	yearpublished, _ := strconv.Atoi(xml.Yearpublished.Value)

	return HotItem{
		ID:            id,
		Rank:          rank,
		Thumbnail:     xml.Thumbnail.Value,
		Name:          xml.Name.Value,
		Yearpublished: yearpublished,
	}
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
		hotItems = append(hotItems, newHotItem(item))
	}

	return HotItems{Items: hotItems}, nil
}
