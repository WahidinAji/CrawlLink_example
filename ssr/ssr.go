package ssr

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"log"
)

type Link struct {
	Url string `json:"url"`
	Title string `json:"title"`
}

func FinishingSsr(url string)  {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	var links []Link
	doc.Find("body a").Each(func(index int, item *goquery.Selection) {
		var link Link
		linkTag := item
		if linkTag.Text() == "\n                Dropdown\n              " {
			link.Title = "Dropdown"
		} else {
			link.Title = linkTag.Text()
		}
		link.Url, _ = item.Attr("href")
		link.Title = link.Title
		//fmt.Println(links)

		links = append(links, link)
	})

	bts, err := json.MarshalIndent(links, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(bts))
}
