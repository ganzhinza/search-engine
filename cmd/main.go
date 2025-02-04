package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"web_crawler/pkg/crawler"
	"web_crawler/pkg/crawler/spider"
	"web_crawler/pkg/index"
	"web_crawler/pkg/netsrv"
)

func main() {
	fileName := "data.txt"
	docks, _ := ScanOrReadDocuments(fileName)
	scanResults := index.New()
	scanResults.AddDocument(docks...)
	saveData(fileName, docks)
	netsrv.ListenAndServe("8000", scanResults)
}

func saveData(fileName string, docks []crawler.Document) {
	if fileName != "" {
		file, err := os.Create(fileName)
		if err != nil {
			log.Println(err)
			return
		}
		defer file.Close()
		pushData(docks, file)
	}
}

func ScanOrReadDocuments(fileName string) ([]crawler.Document, error) {
	var err error
	var docks []crawler.Document
	var file *os.File

	file, err = os.Open(fileName)
	defer file.Close()

	if err != nil {
		var webCrawler crawler.Interface = spider.New()
		docks, err = webCrawler.Scan("https://go.dev", 2)
		if err != nil {
			log.Println(err.Error())
			return nil, fmt.Errorf("Scan error")
		}
	} else {

		docks, err = getData(file)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
	}
	return docks, nil
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
