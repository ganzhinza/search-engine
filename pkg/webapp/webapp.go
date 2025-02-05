package webapp

import (
	"encoding/json"
	"fmt"
	"net/http"

	"web_crawler/pkg/index"

	"github.com/gorilla/mux"
)

var InvIndex index.InvIndex

func indexStorage(w http.ResponseWriter, r *http.Request) {
	buf, err := json.MarshalIndent(InvIndex.Index(), "", "  ")
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

func docsStorage(w http.ResponseWriter, r *http.Request) {
	buf, err := json.MarshalIndent(InvIndex.Documents(), "", "  ")
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

func Web(index index.InvIndex) {
	router := mux.NewRouter()

	InvIndex = index

	endpoints(router)

	http.Handle("/", router)
	http.ListenAndServe("localhost:8080", nil)
}

func endpoints(r *mux.Router) {
	r.HandleFunc("/index", indexStorage)
	r.HandleFunc("/docs", docsStorage)
}
