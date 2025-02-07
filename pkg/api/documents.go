package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"unicode"
	"web_crawler/pkg/crawler"
)

func (api *API) getDocument(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	id, err := strconv.Atoi(params.Get("id"))
	if err != nil {
		fmt.Println(err)
		return
	}

	d := api.index.SearchDocument(id)
	buf, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = w.Write(buf)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func (api *API) deleteDocument(w http.ResponseWriter, r *http.Request) {
	var id int

	params := r.URL.Query()
	id, err := strconv.Atoi(params.Get("id"))
	if err != nil {
		fmt.Println(err)
		return
	}

	api.index.DeleteDocument(id)

	_, err = w.Write([]byte(strconv.Itoa(id)))
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (api *API) updateDocument(w http.ResponseWriter, r *http.Request) {
	var d crawler.Document

	err := json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		fmt.Println(err)
		return
	}

	api.index.UpdateDocument(d)

	_, err = w.Write([]byte("OK"))
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (api *API) addDocument(w http.ResponseWriter, r *http.Request) {
	var d crawler.Document
	err := json.NewDecoder(r.Body).Decode(&d)

	if err != nil {
		fmt.Println(err)
		return
	}

	api.index.AddDocument(d)
}

func (api *API) searchDocuments(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query()
	fullQuery := []rune(param.Get("query"))

	i := 0
	for ; i < len(fullQuery) && unicode.IsLetter(fullQuery[i]); i++ {
	}
	query := string(fullQuery[:i])

	query = strings.ToLower(query)
	queryRes := api.index.GetDocuments(query)

	encoder := json.NewEncoder(w)
	err := encoder.Encode(queryRes)
	if err != nil {
		fmt.Println(err)
		return
	}
}
