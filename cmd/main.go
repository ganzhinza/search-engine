package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"web_crawler/pkg/crawler"
	"web_crawler/pkg/crawler/spider"
	"web_crawler/pkg/index"
)

func main() {
	keyWord := flag.String("s", "Go", "a key word for URLs")
	fileName := flag.String("f", "", "file with scan results")
	flag.Parse()

	var file *os.File
	var err error
	var docks []crawler.Document
	file, err = os.Open(*fileName)
	if err != nil {
		var webCrawler crawler.Interface = spider.New()
		docks, err = webCrawler.Scan("https://go.dev", 2)
		if err != nil {
			log.Println(err.Error())
			return
		}
	} else {
		defer file.Close()
		docks, err = getData(file)
		if err != nil {
			log.Println(err.Error())
			return
		}
	}
	scanResults := index.New()
	for _, doc := range docks {
		scanResults.AddDocument(doc)
	}

	search_res := scanResults.GetDocuments(*keyWord)
	for i := range search_res {
		fmt.Printf("%s\n", search_res[i].Title)
	}

	if *fileName != "" {
		file, err := os.Create(*fileName)
		if err != nil {
			log.Println(err)
			return
		}
		defer file.Close()
		pushData(docks, file)
	}
}

func getData(r io.Reader) ([]crawler.Document, error) {
	buf := make([]byte, 1024*1024)
	bytes := make([]byte, 0, 1024*1024)
	docks := []crawler.Document{}
	num := 0
	var err error
	for num, err = r.Read(buf); num == len(buf) && err == nil; {
		bytes = append(bytes, buf...)
	}
	if err != nil && err != io.EOF {
		return nil, fmt.Errorf("get data %s", err)
	}
	bytes = append(bytes, buf[:num]...)

	err = json.Unmarshal(bytes, &docks)
	if err != nil {
		return nil, fmt.Errorf("get data %s", err)
	}
	return docks, nil
}

func pushData(docks []crawler.Document, w io.Writer) error {
	b, err := json.Marshal(docks)
	if err != nil {
		return fmt.Errorf("push data err: %s", err)
	}

	_, err = w.Write(b)
	if err != nil {
		return fmt.Errorf("push data err: %s", err)
	}
	return nil
}
