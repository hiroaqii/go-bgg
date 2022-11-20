package bgggo

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

type CollectionItemIntermediate struct {
	XMLName    xml.Name `xml:"items"`
	Text       string   `xml:",chardata"`
	Totalitems string   `xml:"totalitems,attr"`
	Termsofuse string   `xml:"termsofuse,attr"`
	Pubdate    string   `xml:"pubdate,attr"`
	Item       []struct {
		Text       string `xml:",chardata"`
		Objecttype string `xml:"objecttype,attr"`
		Objectid   string `xml:"objectid,attr"`
		Subtype    string `xml:"subtype,attr"`
		Collid     string `xml:"collid,attr"`
		Name       struct {
			Text      string `xml:",chardata" json:"text"`
			Sortindex string `xml:"sortindex,attr"`
		} `xml:"name"`
		Yearpublished string `xml:"yearpublished"`
		Image         string `xml:"image"`
		Thumbnail     string `xml:"thumbnail"`
		Stats         struct {
			Text        string `xml:",chardata"`
			Minplayers  string `xml:"minplayers,attr"`
			Maxplayers  string `xml:"maxplayers,attr"`
			Minplaytime string `xml:"minplaytime,attr"`
			Maxplaytime string `xml:"maxplaytime,attr"`
			Playingtime string `xml:"playingtime,attr"`
			Numowned    string `xml:"numowned,attr"`
			Rating      struct {
				Text       string `xml:",chardata"`
				Value      string `xml:"value,attr"`
				Usersrated struct {
					Text  string `xml:",chardata"`
					Value string `xml:"value,attr"`
				} `xml:"usersrated"`
				Average struct {
					Text  string `xml:",chardata"`
					Value string `xml:"value,attr"`
				} `xml:"average"`
				Bayesaverage struct {
					Text  string `xml:",chardata"`
					Value string `xml:"value,attr"`
				} `xml:"bayesaverage"`
				Stddev struct {
					Text  string `xml:",chardata"`
					Value string `xml:"value,attr"`
				} `xml:"stddev"`
				Median struct {
					Text  string `xml:",chardata"`
					Value string `xml:"value,attr"`
				} `xml:"median"`
			} `xml:"rating"`
		} `xml:"stats"`
		Status struct {
			Text             string `xml:",chardata"`
			Own              string `xml:"own,attr"`
			Prevowned        string `xml:"prevowned,attr"`
			Fortrade         string `xml:"fortrade,attr"`
			Want             string `xml:"want,attr"`
			Wanttoplay       string `xml:"wanttoplay,attr"`
			Wanttobuy        string `xml:"wanttobuy,attr"`
			Wishlist         string `xml:"wishlist,attr"`
			Preordered       string `xml:"preordered,attr"`
			Lastmodified     string `xml:"lastmodified,attr"`
			Wishlistpriority string `xml:"wishlistpriority,attr"`
		} `xml:"status"`
		Numplays string `xml:"numplays"`
		Comment  string `xml:"comment"`
	} `xml:"item"`
}

type CollectionItems struct {
	Items []CollectionItem `json:"items"`
}

type CollectionItem struct {
	Name            string `json:"name,omitempty"`
	Yearpublished   int    `json:"yearpublished,omitempty"`
	Image           string `json:"image,omitempty"`
	Thumbnail       string `json:"thumbnail,omitempty"`
	Stats           Stats  `json:"stats,omitempty"`
	Status          Status `json:"status,omitempty"`
	Numplays        int    `json:"numplays,omitempty"`
	Comment         string `json:"comment,omitempty"`
	Conditiontext   string `json:"conditiontext,omitempty"`
	Originalname    string `json:"originalname,omitempty"`
	Wishlistcomment string `json:"wishlistcomment,omitempty"`
}

type Stats struct {
	Minplayers  int    `json:"minplayers,omitempty"`
	Maxplayers  int    `json:"maxplayers,omitempty"`
	Minplaytime int    `json:"minplaytime,omitempty"`
	Maxplaytime int    `json:"maxplaytime,omitempty"`
	Playingtime int    `json:"playingtime,omitempty"`
	Numowned    int    `json:"numowned,omitempty"`
	Rating      Rating `json:"rating,omitempty"`
}

type Rating struct {
	Usersrated   int    `json:"usersrated,omitempty"`
	Average      string `json:"average,omitempty"`
	Bayesaverage string `json:"bayesaverage,omitempty"`
	Stddev       string `json:"stddev,omitempty"`
	Median       string `json:"median,omitempty"`
}

type Status struct {
	Own              int    `json:"own"`
	Prevowned        int    `json:"prevowned"`
	Fortrade         int    `json:"fortrade"`
	Want             int    `json:"want"`
	Wanttoplay       int    `json:"wanttoplay"`
	Wanttobuy        int    `json:"wanttobuy"`
	Wishlist         int    `json:"wishlist"`
	Preordered       int    `json:"preordered"`
	Wishlistpriority int    `json:"wishlistpriority,omitempty"`
	Lastmodified     string `json:"lastmodified"`
}

func Collection() (CollectionItems, error) {
	resp, err := http.Get("https://api.geekdo.com/xmlapi/collection/hiroaqii")
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var itemInter CollectionItemIntermediate
	err = xml.Unmarshal(body, &itemInter)
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	var items = []CollectionItem{}
	for _, item := range itemInter.Item {
		yearpublished, _ := strconv.Atoi(item.Yearpublished)

		usersrated, _ := strconv.Atoi(item.Stats.Rating.Usersrated.Value)
		var rating = Rating{
			Usersrated:   usersrated,
			Average:      item.Stats.Rating.Average.Value,
			Bayesaverage: item.Stats.Rating.Bayesaverage.Value,
			Stddev:       item.Stats.Rating.Stddev.Value,
			Median:       item.Stats.Rating.Median.Value,
		}

		minplayers, _ := strconv.Atoi(item.Stats.Minplayers)
		maxplayers, _ := strconv.Atoi(item.Stats.Maxplayers)
		minplaytime, _ := strconv.Atoi(item.Stats.Minplaytime)
		maxplaytime, _ := strconv.Atoi(item.Stats.Maxplaytime)
		playingtime, _ := strconv.Atoi(item.Stats.Playingtime)
		numowned, _ := strconv.Atoi(item.Stats.Numowned)
		var stats = Stats{
			Rating:      rating,
			Minplayers:  minplayers,
			Maxplayers:  maxplayers,
			Minplaytime: minplaytime,
			Maxplaytime: maxplaytime,
			Playingtime: playingtime,
			Numowned:    numowned,
		}

		own, _ := strconv.Atoi(item.Status.Own)
		prevowned, _ := strconv.Atoi(item.Status.Prevowned)
		fortrade, _ := strconv.Atoi(item.Status.Fortrade)
		want, _ := strconv.Atoi(item.Status.Want)
		wanttoplay, _ := strconv.Atoi(item.Status.Wanttoplay)
		wanttobuy, _ := strconv.Atoi(item.Status.Wanttobuy)
		wishlist, _ := strconv.Atoi(item.Status.Wishlist)
		preordered, _ := strconv.Atoi(item.Status.Preordered)
		wishlistpriority, _ := strconv.Atoi(item.Status.Wishlistpriority)

		var status = Status{
			Own:              own,
			Prevowned:        prevowned,
			Fortrade:         fortrade,
			Want:             want,
			Wanttoplay:       wanttoplay,
			Wanttobuy:        wanttobuy,
			Wishlist:         wishlist,
			Preordered:       preordered,
			Wishlistpriority: wishlistpriority,
			Lastmodified:     item.Status.Lastmodified,
		}
		items = append(items, CollectionItem{
			Name:          item.Name.Text,
			Yearpublished: yearpublished,
			Comment:       item.Comment,
			//Thumbnail:     item.Thumbnail,
			Stats:  stats,
			Status: status,
		})
	}

	return CollectionItems{Items: items}, nil
}
