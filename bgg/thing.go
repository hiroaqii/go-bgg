package bgg

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

type thingIntermediate struct {
	XMLName    xml.Name                `xml:"items"`
	Text       string                  `xml:",chardata"`
	Termsofuse string                  `xml:"termsofuse,attr"`
	Item       []thingItemIntermediate `xml:"item"`
}

type thingItemIntermediate struct {
	Text      string `xml:",chardata"`
	Type      string `xml:"type,attr"`
	ID        string `xml:"id,attr"`
	Thumbnail string `xml:"thumbnail"`
	Image     string `xml:"image"`
	Name      []struct {
		Text      string `xml:",chardata"`
		Type      string `xml:"type,attr"`
		Sortindex string `xml:"sortindex,attr"`
		Value     string `xml:"value,attr"`
	} `xml:"name"`
	Description   string `xml:"description"`
	Yearpublished struct {
		Text  string `xml:",chardata"`
		Value string `xml:"value,attr"`
	} `xml:"yearpublished"`
	Minplayers struct {
		Text  string `xml:",chardata"`
		Value string `xml:"value,attr"`
	} `xml:"minplayers"`
	Maxplayers struct {
		Text  string `xml:",chardata"`
		Value string `xml:"value,attr"`
	} `xml:"maxplayers"`
	Playingtime struct {
		Text  string `xml:",chardata"`
		Value string `xml:"value,attr"`
	} `xml:"playingtime"`
	Minplaytime struct {
		Text  string `xml:",chardata"`
		Value string `xml:"value,attr"`
	} `xml:"minplaytime"`
	Maxplaytime struct {
		Text  string `xml:",chardata"`
		Value string `xml:"value,attr"`
	} `xml:"maxplaytime"`
	Minage struct {
		Text  string `xml:",chardata"`
		Value string `xml:"value,attr"`
	} `xml:"minage"`
	Poll []thingPoolIntermediate `xml:"poll"`
	Link []thingLinkIntermediate `xml:"link"`
}
type thingPoolIntermediate struct {
	Text       string `xml:",chardata"`
	Name       string `xml:"name,attr"`
	Title      string `xml:"title,attr"`
	Totalvotes string `xml:"totalvotes,attr"`
	Results    []struct {
		Text       string `xml:",chardata"`
		Numplayers string `xml:"numplayers,attr"`
		Result     []struct {
			Text     string `xml:",chardata"`
			Value    string `xml:"value,attr"`
			Numvotes string `xml:"numvotes,attr"`
			Level    string `xml:"level,attr"`
		} `xml:"result"`
	} `xml:"results"`
}
type thingLinkIntermediate struct {
	Text  string `xml:",chardata"`
	Type  string `xml:"type,attr"`
	ID    string `xml:"id,attr"`
	Value string `xml:"value,attr"`
}

type ThingItems struct {
	Items []ThingItem `json:"items"`
}

type ThingItem struct {
	ID            int         `json:"id"`
	Thumbnail     string      `json:"thumbnail,omitempty"`
	Image         string      `json:"image,omitempty"`
	Name          string      `json:"name,omitempty"`
	Description   string      `json:"description,omitempty"`
	Yearpublished int         `json:"yearpublished"`
	Minplayers    int         `json:"minplayers,omitempty"`
	Maxplayers    int         `json:"maxplayers,omitempty"`
	Minplaytime   int         `json:"minplaytime,omitempty"`
	Maxplaytime   int         `json:"maxplaytime,omitempty"`
	Minage        int         `json:"minage,omitempty"`
	Playingtime   int         `json:"playingtime,omitempty"`
	Polls         []ThingPool `json:"poll,omitempty"`
	Links         []ThingLink `json:"link,omitempty"`
}

type ThingPool struct {
	Name       string             `json:"name,omitempty"`
	Title      string             `json:"title,omitempty"`
	Totalvotes int                `json:"totalvotes"`
	Results    []ThingPoolResults `json:"results,omitempty"`
}

type ThingPoolResults struct {
	Numplayers string                   `json:"numplayers,omitempty"`
	Result     []ThingPoolResultsResult `json:"result,omitempty"`
}

type ThingPoolResultsResult struct {
	Level    int    `json:"level,omitempty"`
	Value    string `json:"value,omitempty"`
	Numvotes int    `json:"numvotes"`
}

type ThingLink struct {
	ID    int    `json:"id,omitempty"`
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}

func newThingItem(item thingItemIntermediate) ThingItem {
	id, _ := strconv.Atoi(item.ID)
	yearpublished, _ := strconv.Atoi(item.Yearpublished.Value)
	minplayers, _ := strconv.Atoi(item.Minplayers.Value)
	maxplayers, _ := strconv.Atoi(item.Maxplayers.Value)
	minplaytime, _ := strconv.Atoi(item.Minplaytime.Value)
	maxplaytime, _ := strconv.Atoi(item.Maxplaytime.Value)
	minage, _ := strconv.Atoi(item.Minage.Value)
	playingtime, _ := strconv.Atoi(item.Playingtime.Value)

	name := ""
	for _, nm := range item.Name {
		if nm.Type == "primary" {
			name = nm.Value
		}
	}

	polls := []ThingPool{}
	for _, pool := range item.Poll {
		polls = append(polls, newThingPool(pool))
	}

	links := []ThingLink{}
	for _, link := range item.Link {
		links = append(links, newThingLink(link))
	}

	return ThingItem{
		ID:            id,
		Name:          name,
		Thumbnail:     item.Thumbnail,
		Image:         item.Image,
		Description:   item.Description,
		Yearpublished: yearpublished,
		Minplayers:    minplayers,
		Maxplayers:    maxplayers,
		Minplaytime:   minplaytime,
		Maxplaytime:   maxplaytime,
		Minage:        minage,
		Playingtime:   playingtime,
		Polls:         polls,
		Links:         links,
	}
}

func newThingPool(pool thingPoolIntermediate) ThingPool {
	var poolResults = []ThingPoolResults{}
	for _, results := range pool.Results {

		var resultsResult = []ThingPoolResultsResult{}
		for _, r := range results.Result {
			level, _ := strconv.Atoi(r.Level)
			numvotes, _ := strconv.Atoi(r.Numvotes)
			resultsResult = append(resultsResult, ThingPoolResultsResult{
				Level:    level,
				Value:    r.Value,
				Numvotes: numvotes,
			})
		}

		poolResults = append(poolResults, ThingPoolResults{
			Numplayers: results.Numplayers,
			Result:     resultsResult,
		})
	}

	totalvotes, _ := strconv.Atoi(pool.Totalvotes)
	return ThingPool{
		Name:       pool.Name,
		Title:      pool.Title,
		Totalvotes: totalvotes,
		Results:    poolResults,
	}
}

func newThingLink(t thingLinkIntermediate) ThingLink {
	id, _ := strconv.Atoi(t.ID)
	return ThingLink{
		ID:    id,
		Type:  t.Type,
		Value: t.Value,
	}
}

func Thing(id string) (ThingItems, error) {
	resp, err := http.Get("https://api.geekdo.com/xmlapi2/thing?id=" + id)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var itemInter thingIntermediate
	err = xml.Unmarshal(body, &itemInter)
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	var thingItems = []ThingItem{}
	for _, item := range itemInter.Item {
		thingItems = append(thingItems, newThingItem(item))
	}

	return ThingItems{Items: thingItems}, nil
}
