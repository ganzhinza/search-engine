package webapp

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"web_crawler/pkg/crawler"
	"web_crawler/pkg/index"

	"github.com/gorilla/mux"
)

var testMux *mux.Router

func TestMain(m *testing.M) {
	testMux = mux.NewRouter()
	endpoints(testMux)
	InvIndex = *index.New()
	InvIndex.AddDocument(crawler.Document{ID: 0, URL: "0", Title: "0", Body: "0"},
		crawler.Document{ID: 1, URL: "1", Title: "1", Body: "1"},
		crawler.Document{ID: 2, URL: "2", Title: "2", Body: "2"},
		crawler.Document{ID: 3, URL: "3", Title: "3", Body: "3"})
	m.Run()
}

func Test_docsStorage(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/docs", nil)
	req.Header.Add("Сontent-type", "plain/text")

	rr := httptest.NewRecorder()

	testMux.ServeHTTP(rr, req)

	body := rr.Body.String()
	for i := 0; i < 4; i++ {
		if !strings.Contains(body, strconv.Itoa(i)) {
			t.Fatal(body)
		}
	}
}

func Test_indexStorage(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/index", nil)

	rr := httptest.NewRecorder()

	testMux.ServeHTTP(rr, req)

	body := rr.Body.String()
	for i := 0; i < 4; i++ {
		if !strings.Contains(body, strconv.Itoa(i)) {
			t.Fatal(body)
		}
	}

}
