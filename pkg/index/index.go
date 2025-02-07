package index

import (
	"strings"
	"unicode"
	"web_crawler/pkg/crawler"
)

type InvIndex struct {
	documents []crawler.Document
	index     map[string][]int
}

func New() *InvIndex {
	index := InvIndex{}
	index.index = make(map[string][]int)
	return &index
}

func (i *InvIndex) Index() map[string][]int {
	return i.index
}

func (i *InvIndex) Documents() []crawler.Document {
	return i.documents
}

func (index *InvIndex) AddDocument(docks ...crawler.Document) {
	for _, dock := range docks {
		dock.ID = len(index.documents)
		index.documents = append(index.documents, dock)
		words := make(map[string]bool)

		pureWords := make([]string, 0, 10)
		start := 0
		for end, l := range dock.Title {
			if !unicode.IsLetter(l) {
				keyWord := strings.ToLower(dock.Title[start:end])
				if end-start > 0 && !words[keyWord] {
					pureWords = append(pureWords, keyWord)
					words[keyWord] = true
				}
				start = end + 1
			}
		}
		keyWord := strings.ToLower(dock.Title[start:])
		if start != len(dock.Title)-1 {
			pureWords = append(pureWords, keyWord)
		}

		for _, word := range pureWords {
			index.index[word] = append(index.index[word], dock.ID)
		}
	}

}

func (index *InvIndex) GetDocuments(word string) []crawler.Document {
	docs := make([]crawler.Document, 0, 10)
	for _, id := range index.index[word] {
		docs = append(docs, index.SearchDocument(id))
	}
	return docs
}

func (index *InvIndex) searchDocument(id int) int {
	l := 0
	r := len(index.documents)
	m := (l + r) / 2
	for l <= r {
		m = l + (r-l)/2
		if id == index.documents[m].ID {
			return m
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
	return -1
}

func (index *InvIndex) SearchDocument(id int) crawler.Document {
	return index.documents[index.searchDocument(id)]
}

func (index *InvIndex) DeleteDocument(id int) bool {
	i := index.searchDocument(id)
	if i == -1 {
		return false
	} else {
		if i == len(index.documents) {
			index.documents = index.documents[:len(index.documents)-1]
		} else {
			index.documents = append(index.documents[:i], index.documents[i+1:]...)
		}
		return true
	}
}

func (index *InvIndex) UpdateDocument(doc crawler.Document) bool {
	i := index.searchDocument(doc.ID)
	if i == -1 {
		return false
	} else {
		index.documents[i] = doc
		return true
	}
}
