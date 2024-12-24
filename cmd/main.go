package main

import (
	"flag"
	"fmt"
	"web_crawler/pkg/crawler"
	"web_crawler/pkg/crawler/spider"
	"web_crawler/pkg/index"
)

func main() {
	keyWord := flag.String("s", "Go", "a key word for URLs")
	flag.Parse()

	var webCrawler crawler.Interface = spider.New()

	scanResults := index.New()

	docksDev, err := webCrawler.Scan("https://go.dev", 2)
	if err != nil {
		fmt.Printf("%s", err)
	}

	for _, doc := range docksDev {
		scanResults.AddDocument(doc)
	}

	search_res := scanResults.GetDocuments(*keyWord) //i don't know *is much lighter but obj is musch safer
	for i := range search_res {
		fmt.Printf("%s\n", search_res[i].Title)
	}

}
