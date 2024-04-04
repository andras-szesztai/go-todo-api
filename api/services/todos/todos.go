package todos

import (
	"net/http"
	"todo-api/db"
	"todo-api/utils"

	"github.com/gorilla/mux"
)

type TodosService struct {
	store db.Store
}

func NewTodosService(store db.Store) *TodosService {
	return &TodosService{store: store}
}

func (ts *TodosService) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/todos", ts.handleGetTodos).Methods("GET")
}

func (ts *TodosService) handleGetTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := ts.store.GetTodos()
	if err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJson(w, http.StatusOK, todos)
}
