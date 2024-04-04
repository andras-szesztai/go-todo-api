package api

import (
	"log"
	"net/http"
	"todo-api/db"
	"todo-api/services/todos"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr  string
	store db.Store
}

func NewAPIServer(addr string, store db.Store) *APIServer {
	return &APIServer{addr: addr, store: store}
}

func (server *APIServer) Start() {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	taskService := todos.NewTodosService(server.store)
	taskService.RegisterRoutes(subRouter)

	log.Println("Starting API server on", server.addr)

	log.Fatal(http.ListenAndServe(server.addr, subRouter))
}
