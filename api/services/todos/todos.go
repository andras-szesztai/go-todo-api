package todos

import (
	"encoding/json"
	"errors"
	"io"
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
	r.HandleFunc("/todos", ts.handlePostTodo).Methods(http.MethodPost)
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

func (ts *TodosService) handlePostTodo(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, err)
		return
	}
	defer r.Body.Close()
	var todo *types.PostTodoRequest
	if err = json.Unmarshal(body, &todo); err != nil {
		utils.WriteJson(w, http.StatusBadRequest, types.ErrorResponse{Error: "Invalid request payload"})
		return
	}

	if err := validateTodoPayload(todo); err != nil {
		utils.WriteJson(w, http.StatusBadRequest, types.ErrorResponse{Error: err.Error()})
		return
	}

	if createdTodo, err := ts.store.PostTodo(todo); err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, types.ErrorResponse{Error: "Error creating task"})
		return
	} else {
		utils.WriteJson(w, http.StatusCreated, createdTodo)
	}
}

func validateTodoPayload(todo *types.PostTodoRequest) error {
	if todo.Name == "" {
		return errors.New("todo name is required")
	}
	return nil
}
