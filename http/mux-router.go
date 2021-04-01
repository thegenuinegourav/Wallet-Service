package http

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"sync"
)

type muxRouter struct {
	client *mux.Router
}

var (
	mR 			*muxRouter
	routerOnce 	sync.Once
)

// Making muxRouter instance as singleton
func NewMuxRouter() IRouter {
	if mR == nil {
		routerOnce.Do(func() {
			client := mux.NewRouter().StrictSlash(true)
			mR = &muxRouter{client}
		})
	}
	return mR
}

func (r *muxRouter) ADDVERSION(uri string) {
	r.client = r.client.PathPrefix(uri).Subrouter()
}

func (r *muxRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	r.client.HandleFunc(uri,f).Methods("GET")
}

func (r *muxRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	r.client.HandleFunc(uri,f).Methods("POST")
}

func (r *muxRouter) PUT(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	r.client.HandleFunc(uri,f).Methods("PUT")
}

func (r *muxRouter) DELETE(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	r.client.HandleFunc(uri,f).Methods("DELETE")
}

func (r *muxRouter) SERVE(port string) {
	log.Printf("Mux HTTP server running on port %v", port)
	http.ListenAndServe(":" + port, r.client)
}
