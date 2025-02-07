package api

import (
	"net/http"
	"web_crawler/pkg/index"

	"github.com/gorilla/mux"
)

type API struct {
	router *mux.Router
	index  *index.InvIndex
}

func New(index *index.InvIndex) *API {
	api := API{
		router: mux.NewRouter(),
		index:  index,
	}

	api.endpoints()
	return &api
}

// CRUD
func (api *API) endpoints() {
	api.router.HandleFunc("/api/v1/getSearchResult", api.searchDocuments).Methods(http.MethodGet)
	api.router.HandleFunc("/api/v1/getDocument", api.getDocument).Methods(http.MethodGet)
	api.router.HandleFunc("/api/v1/deleteDocument", api.deleteDocument).Methods(http.MethodDelete)
	api.router.HandleFunc("/api/v1/updateDocument", api.updateDocument).Methods(http.MethodPut)
	api.router.HandleFunc("/api/v1/addDocument", api.addDocument).Methods(http.MethodPost)
}

func (api *API) Router() *mux.Router {
	return api.router
}
