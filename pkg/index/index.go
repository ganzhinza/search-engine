package index

import (
	"unicode"
	"web_crawler/pkg/crawler"
)

type InvIndex struct {
	documents []crawler.Document
	Index     map[string][]int
}

func New() *InvIndex {
	index := InvIndex{}
	index.Index = make(map[string][]int)
	return &index
}

func (index *InvIndex) AddDocument(document crawler.Document) {
	document.ID = len(index.documents)
	index.documents = append(index.documents, document)

	pureWords := make([]string, 0, 10)
	start := 0
	for end, l := range document.Title {
		if !unicode.IsLetter(l) {
			if end-start > 0 {
				pureWords = append(pureWords, document.Title[start:end])
			}
			start = end + 1
		}
	}
	if start != len(document.Title)-1 {
		pureWords = append(pureWords, document.Title[start:])
	}

	for _, word := range pureWords {
		index.Index[word] = append(index.Index[word], document.ID)
	}
}

func (index *InvIndex) GetDocuments(word string) []crawler.Document {
	docs := make([]crawler.Document, 0, 10)
	for _, id := range index.Index[word] {
		docs = append(docs, index.SearchDocument(id))
	}
	return docs
}

func (index *InvIndex) SearchDocument(id int) crawler.Document {
	l := 0
	r := len(index.documents)
	m := (l + r) / 2
	for l <= r {
		m = l + (r-l)/2
		if id == index.documents[m].ID {
			return index.documents[m]
		}
		if id > index.documents[m].ID {
			l = m + 1
			continue
		}
		if id < index.documents[m].ID {
			r = m - 1
			continue
		}
	}
	return crawler.Document{ID: -1}
}

func (index *InvIndex) GetAllDocs() []crawler.Document {
	return index.documents
}

func (index *InvIndex) GetAllWords() []string {
	words := make([]string, 0, len(index.Index))
	for word := range index.Index {
		words = append(words, word)
	}
	return words
}
