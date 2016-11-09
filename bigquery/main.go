package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/PuerkitoBio/goquery"

	"github.com/gotokatsuya/incidents/slack"
)

var (
	incidentID int
	slackURL   string

	cacheIncidentInfoMap map[string]struct{}
)

func init() {
	flag.IntVar(&incidentID, "i", 18022, "incident")
	flag.StringVar(&slackURL, "u", "", "slack url")
	cacheIncidentInfoMap = make(map[string]struct{})
}

func postFetchedIncidentInfo(url string) {
	doc, _ := goquery.NewDocument(url)
	doc.Find("table > tbody > tr > td").Each(func(_ int, s *goquery.Selection) {
		msg := s.Text()
		if _, ok := cacheIncidentInfoMap[msg]; ok {
			return
		}
		fmt.Println(msg)
		if len(slackURL) != 0 {
			slack.Post(slack.Data{
				Text:      msg,
				Username:  "bigquery incidents",
				IconEmoji: ":see_no_evil:",
			}, slackURL)
		}
		cacheIncidentInfoMap[msg] = struct{}{}
	})
}

func main() {
	url := fmt.Sprintf("https://status.cloud.google.com/incident/bigquery/%d", incidentID)

	for {
		postFetchedIncidentInfo(url)

		time.Sleep(1 * time.Minute)
	}
}
