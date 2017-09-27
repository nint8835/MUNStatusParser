package main

import(
	"github.com/grokify/html-strip-tags-go"
	"strings"
	"encoding/json"
	"net/http"
	"log"
	"io/ioutil"
)

const MUN_URL string = "https://mun.apparmor.com/notificationhistory/"

type FeedItem struct{
	DescriptionHTML string
	IconCharCode    string
	State           string
	Title           string
}

func (i FeedItem) Description() string{
	newline_replaced := strings.Replace(i.DescriptionHTML, "<br>", "\n", -1)
	tags_fixed := strings.Replace(newline_replaced, "\\/", "/", -1)
	return strip.StripTags(tags_fixed)
}

type Feed struct{
	FeedItems []FeedItem
}

func GetFeed() Feed{
	return GetFeedFromUrl(MUN_URL)
}

func GetFeedFromUrl(url string) Feed{
	response, err := http.Get(url)
	if err != nil{
		log.Fatal(err)
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return Parse(data)

}

func Parse(b []byte) Feed{
	var feed Feed
	err := json.Unmarshal(b, &feed)
	if err != nil{
		log.Fatal(err)
	}
	return feed
}