package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"web_crawler/pkg/crawler"
	"web_crawler/pkg/index"

	"github.com/gorilla/mux"
)

var api API

func TestMain(m *testing.M) {
	InvIndex := index.New()
	InvIndex.AddDocument(crawler.Document{ID: 0, URL: "0", Title: "0", Body: "0"},
		crawler.Document{ID: 1, URL: "1", Title: "1", Body: "1"},
		crawler.Document{ID: 2, URL: "2", Title: "2", Body: "2"},
		crawler.Document{ID: 3, URL: "3", Title: "3", Body: "3"})

	api = API{
		router: mux.NewRouter(),
		index:  InvIndex,
	}

	api.endpoints()
	m.Run()
}

func TestAPI_searchDocument(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/v1/getSearchResult?query=0", nil)
	want := crawler.Document{ID: 0, URL: "0", Title: "0", Body: "0"}
	req.Header.Add("Сontent-type", "plain/text")

	rr := httptest.NewRecorder()

	api.router.ServeHTTP(rr, req)

	d := []crawler.Document{}
	decoder := json.NewDecoder(rr.Body)
	decoder.Decode(&d)

	if d[0] != want {
		t.Fatal(d)
		return
	}
}

func TestAPI_addDocument(t *testing.T) {
	inputDocument := crawler.Document{ID: 0, URL: "0", Title: "0", Body: "0"}
	jsonWant, err := json.Marshal(inputDocument)
	if err != nil {
		t.Fatal(err)
		return
	}

	req := httptest.NewRequest(http.MethodPost, "/api/v1/addDocument", bytes.NewBuffer(jsonWant))
	req.Header.Add("Сontent-type", "plain/text")

	rr := httptest.NewRecorder()

	api.router.ServeHTTP(rr, req)
	docs := api.index.Documents()
	inputDocument.ID = len(docs) - 1
	if inputDocument != docs[len(docs)-1] {
		t.Fatal(inputDocument, docs[len(docs)-1])
		return
	}

}

func TestAPI_deleteDocument(t *testing.T) {
	id := 1
	req := httptest.NewRequest(http.MethodDelete, "/api/v1/deleteDocument?id="+strconv.Itoa(id), nil)
	req.Header.Add("Сontent-type", "plain/text")

	rr := httptest.NewRecorder()

	api.router.ServeHTTP(rr, req)

	docs := api.index.Documents()

	if id == docs[id].ID {
		t.Fatal("not deleted")
		return
	}

}

func TestAPI_getDocument(t *testing.T) {
	id := 1
	want := api.index.SearchDocument(id)
	got := crawler.Document{}

	req := httptest.NewRequest(http.MethodGet, "/api/v1/getDocument?id="+strconv.Itoa(id), nil)
	req.Header.Add("Сontent-type", "plain/text")

	rr := httptest.NewRecorder()

	api.router.ServeHTTP(rr, req)

	decoder := json.NewDecoder(rr.Body)
	err := decoder.Decode(&got)
	if err != nil {
		t.Fatal(err)
	}

	if got != want {
		t.Fatal("Unexpected result got: ", got, ", want: ", want)
	}
}

func TestAPI_updateDocument(t *testing.T) {
	want := crawler.Document{ID: 0, URL: "HOHOHO", Title: "LOL"}
	var got crawler.Document
	buf, err := json.Marshal(want)
	if err != nil {
		t.Fatal(err)
		return
	}

	req := httptest.NewRequest(http.MethodPut, "/api/v1/updateDocument", bytes.NewBuffer(buf))

	req.Header.Add("Сontent-type", "plain/text")

	rr := httptest.NewRecorder()
	api.router.ServeHTTP(rr, req)

	got = api.index.SearchDocument(want.ID)

	if got != want {
		t.Fatal("Unexpected result got: ", got, ", want: ", want)
	}

}
