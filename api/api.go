package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr  string
	store Store
}

func NewAPIServer(addr string, store Store) *APIServer {
	return &APIServer{addr: addr, store: store}
}

func (server *APIServer) Start() {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	subRouter.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	}).Methods("GET")

	log.Println("Starting API server on", server.addr)

	log.Fatal(http.ListenAndServe(server.addr, subRouter))
}
