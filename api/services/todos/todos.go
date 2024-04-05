package todos

import (
	"net/http"
	"todo-api/db"
	"todo-api/types"
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
	r.HandleFunc("/todos", ts.handleGetTodos).Methods(http.MethodGet)
	r.HandleFunc("/todos/{id}", ts.handleGetTodoById).Methods(http.MethodGet)
}

func (ts *TodosService) handleGetTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := ts.store.GetTodos()
	if err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJson(w, http.StatusOK, todos)
}

func (ts *TodosService) handleGetTodoById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	todo, err := ts.store.GetTodoById(id)
	if err != nil {
		utils.WriteJson(w, http.StatusNotFound, types.ErrorResponse{Error: "todo not found"})
		return
	}
	utils.WriteJson(w, http.StatusOK, todo)
}
