package bgg

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

type collectionItemsIntermediate struct {
	XMLName    xml.Name                     `xml:"items"`
	Text       string                       `xml:",chardata"`
	Totalitems string                       `xml:"totalitems,attr"`
	Termsofuse string                       `xml:"termsofuse,attr"`
	Pubdate    string                       `xml:"pubdate,attr"`
	Item       []collectionItemIntermediate `xml:"item"`
}

type collectionItemIntermediate struct {
	Text       string `xml:",chardata"`
	Objecttype string `xml:"objecttype,attr"`
	Objectid   string `xml:"objectid,attr"`
	Subtype    string `xml:"subtype,attr"`
	Collid     string `xml:"collid,attr"`
	Name       struct {
		Text      string `xml:",chardata" json:"text"`
		Sortindex string `xml:"sortindex,attr"`
	} `xml:"name"`
	Yearpublished string                           `xml:"yearpublished"`
	Image         string                           `xml:"image"`
	Thumbnail     string                           `xml:"thumbnail"`
	Stats         collectionItemStatsIntermediate  `xml:"stats"`
	Status        collectionItemStatusIntermediate `xml:"status"`
	Numplays      string                           `xml:"numplays"`
	Comment       string                           `xml:"comment"`
}

type collectionItemStatsIntermediate struct {
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
}
type collectionItemStatusIntermediate struct {
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
}

type CollectionItems struct {
	Items []CollectionItem `json:"items"`
}

type CollectionItem struct {
	Name            string               `json:"name,omitempty"`
	Yearpublished   int                  `json:"yearpublished,omitempty"`
	Image           string               `json:"image,omitempty"`
	Thumbnail       string               `json:"thumbnail,omitempty"`
	Stats           CollectionItemStats  `json:"stats,omitempty"`
	Status          CollectionItemStatus `json:"status,omitempty"`
	Numplays        int                  `json:"numplays,omitempty"`
	Comment         string               `json:"comment,omitempty"`
	Conditiontext   string               `json:"conditiontext,omitempty"`
	Originalname    string               `json:"originalname,omitempty"`
	Wishlistcomment string               `json:"wishlistcomment,omitempty"`
}

type CollectionItemStats struct {
	Minplayers  int                  `json:"minplayers,omitempty"`
	Maxplayers  int                  `json:"maxplayers,omitempty"`
	Minplaytime int                  `json:"minplaytime,omitempty"`
	Maxplaytime int                  `json:"maxplaytime,omitempty"`
	Playingtime int                  `json:"playingtime,omitempty"`
	Numowned    int                  `json:"numowned,omitempty"`
	Rating      CollectionItemRating `json:"rating,omitempty"`
}

type CollectionItemRating struct {
	Usersrated   int    `json:"usersrated,omitempty"`
	Average      string `json:"average,omitempty"`
	Bayesaverage string `json:"bayesaverage,omitempty"`
	Stddev       string `json:"stddev,omitempty"`
	Median       string `json:"median,omitempty"`
}

type CollectionItemStatus struct {
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

func newCollectionItemStats(s collectionItemStatsIntermediate) CollectionItemStats {
	usersrated, _ := strconv.Atoi(s.Rating.Usersrated.Value)
	var rating = CollectionItemRating{
		Usersrated:   usersrated,
		Average:      s.Rating.Average.Value,
		Bayesaverage: s.Rating.Bayesaverage.Value,
		Stddev:       s.Rating.Stddev.Value,
		Median:       s.Rating.Median.Value,
	}

	minplayers, _ := strconv.Atoi(s.Minplayers)
	maxplayers, _ := strconv.Atoi(s.Maxplayers)
	minplaytime, _ := strconv.Atoi(s.Minplaytime)
	maxplaytime, _ := strconv.Atoi(s.Maxplaytime)
	playingtime, _ := strconv.Atoi(s.Playingtime)
	numowned, _ := strconv.Atoi(s.Numowned)

	return CollectionItemStats{
		Rating:      rating,
		Minplayers:  minplayers,
		Maxplayers:  maxplayers,
		Minplaytime: minplaytime,
		Maxplaytime: maxplaytime,
		Playingtime: playingtime,
		Numowned:    numowned,
	}
}

func newCollectionItemStatus(s collectionItemStatusIntermediate) CollectionItemStatus {
	own, _ := strconv.Atoi(s.Own)
	prevowned, _ := strconv.Atoi(s.Prevowned)
	fortrade, _ := strconv.Atoi(s.Fortrade)
	want, _ := strconv.Atoi(s.Want)
	wanttoplay, _ := strconv.Atoi(s.Wanttoplay)
	wanttobuy, _ := strconv.Atoi(s.Wanttobuy)
	wishlist, _ := strconv.Atoi(s.Wishlist)
	preordered, _ := strconv.Atoi(s.Preordered)
	wishlistpriority, _ := strconv.Atoi(s.Wishlistpriority)

	return CollectionItemStatus{
		Own:              own,
		Prevowned:        prevowned,
		Fortrade:         fortrade,
		Want:             want,
		Wanttoplay:       wanttoplay,
		Wanttobuy:        wanttobuy,
		Wishlist:         wishlist,
		Preordered:       preordered,
		Wishlistpriority: wishlistpriority,
		Lastmodified:     s.Lastmodified,
	}
}

func newCollectionItem(item collectionItemIntermediate) CollectionItem {
	stats := newCollectionItemStats(item.Stats)
	status := newCollectionItemStatus(item.Status)
	yearpublished, _ := strconv.Atoi(item.Yearpublished)

	return CollectionItem{
		Name:          item.Name.Text,
		Yearpublished: yearpublished,
		Comment:       item.Comment,
		Thumbnail:     item.Thumbnail,
		Stats:         stats,
		Status:        status,
	}
}

func Collection(username string) (CollectionItems, error) {
	resp, err := http.Get("https://api.geekdo.com/xmlapi/collection/" + username)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var itemInter collectionItemsIntermediate
	err = xml.Unmarshal(body, &itemInter)
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	var items = []CollectionItem{}
	for _, item := range itemInter.Item {
		items = append(items, newCollectionItem(item))
	}

	return CollectionItems{Items: items}, nil
}
