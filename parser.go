package munstatusparser

import(
	"github.com/grokify/html-strip-tags-go"
	"strings"
	"encoding/json"
	"net/http"
	"io/ioutil"
	"regexp"
)

const MUN_URL string = "https://mun.apparmor.com/notificationhistory/"

type FeedItem struct{
	DescriptionHTML string
	IconCharCode    string
	State           string
	Title           string
}

func (i FeedItem) CleanText() string{
	newlineReplaced := strings.Replace(i.DescriptionHTML, "<br>", "\n", -1)
	tagsFixed := strings.Replace(newlineReplaced, "\\/", "/", -1)
	return strings.TrimSpace(strip.StripTags(tagsFixed))
}

func (i FeedItem) Description() string{
	regex := regexp.MustCompile(`(?m)^.+\n([\S \n]+)\n.+$`)
	return strings.TrimSpace(regex.FindStringSubmatch(i.CleanText())[1])
}

func (i FeedItem) SentTime() string{
	regex := regexp.MustCompile(`(?m)^[\S \n]+\n(.+)$`)
	return strings.TrimSpace(regex.FindStringSubmatch(i.CleanText())[1])
}

type Feed struct{
	FeedItems []FeedItem
}

func GetFeed() (Feed, error){
	return GetFeedFromUrl(MUN_URL)
}

func GetFeedFromUrl(url string) (Feed, error){
	response, err := http.Get(url)
	if err != nil{
		return Feed{}, nil
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return Feed{}, nil
	}

	return Parse(data)

}

func Parse(b []byte) (Feed, error){
	var feed Feed
	err := json.Unmarshal(b, &feed)
	if err != nil{
		return Feed{}, err
	}
	return feed, nil
}