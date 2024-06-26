package todos

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo-api/types"
	"todo-api/utils"

	"github.com/gorilla/mux"
)

func TestGetTodos(t *testing.T) {
	t.Run("should return todos on success", func(t *testing.T) {
		ms := &utils.MockStore{}
		service := NewTodosService(ms)
		req, error := http.NewRequest(http.MethodGet, "/todos", nil)
		if error != nil {
			t.Fatal(error)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/todos", service.handleGetTodos)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status code %d, got %d", http.StatusOK, rr.Code)
		}
	})
	t.Run("should return 500 on error", func(t *testing.T) {
		ms := &utils.MockErrorStore{}
		service := NewTodosService(ms)
		req, error := http.NewRequest(http.MethodGet, "/todos", nil)
		if error != nil {
			t.Fatal(error)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/todos", service.handleGetTodos)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusInternalServerError {
			t.Errorf("expected status code %d, got %d", http.StatusInternalServerError, rr.Code)
		}
	})
}

func TestGetTodoById(t *testing.T) {
	t.Run("should return todo on success", func(t *testing.T) {
		ms := &utils.MockStore{}
		service := NewTodosService(ms)
		req, error := http.NewRequest(http.MethodGet, "/todos/1", nil)
		if error != nil {
			t.Fatal(error)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/todos/{id}", service.handleGetTodoById)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status code %d, got %d", http.StatusOK, rr.Code)
		}
	})
	t.Run("should return 404 if todo not found", func(t *testing.T) {
		ms := &utils.MockErrorStore{}
		service := NewTodosService(ms)
		req, error := http.NewRequest(http.MethodGet, "/todos/not-existing-id", nil)
		if error != nil {
			t.Fatal(error)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/todos/{id}", service.handleGetTodoById)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusNotFound {
			t.Errorf("expected status code %d, got %d", http.StatusNotFound, rr.Code)
		}
	})
}

func TestPostTodo(t *testing.T) {
	t.Run("should create todo on success", func(t *testing.T) {
		payload := &types.PostTodoRequest{
			Name: "Test todo",
		}

		b, error := json.Marshal(payload)
		if error != nil {
			t.Fatal(error)
		}

		ms := &utils.MockStore{}
		service := NewTodosService(ms)
		req, error := http.NewRequest(http.MethodPost, "/todos", bytes.NewBuffer(b))
		if error != nil {
			t.Fatal(error)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/todos", service.handlePostTodo)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf("expected status code %d, got %d", http.StatusCreated, rr.Code)
		}
	})

	t.Run("should return 400 on invalid payload", func(t *testing.T) {
		payload := &types.PostTodoRequest{
			Name: "",
		}

		b, error := json.Marshal(payload)
		if error != nil {
			t.Fatal(error)
		}

		ms := &utils.MockStore{}
		service := NewTodosService(ms)
		req, error := http.NewRequest(http.MethodPost, "/todos", bytes.NewBuffer(b))
		if error != nil {
			t.Fatal(error)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/todos", service.handlePostTodo)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("should return 500 on error", func(t *testing.T) {
		payload := &types.PostTodoRequest{
			Name: "Test todo",
		}

		b, error := json.Marshal(payload)
		if error != nil {
			t.Fatal(error)
		}

		ms := &utils.MockErrorStore{}
		service := NewTodosService(ms)
		req, error := http.NewRequest(http.MethodPost, "/todos", bytes.NewBuffer(b))
		if error != nil {
			t.Fatal(error)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/todos", service.handlePostTodo)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusInternalServerError {
			t.Errorf("expected status code %d, got %d", http.StatusInternalServerError, rr.Code)
		}
	})
}
