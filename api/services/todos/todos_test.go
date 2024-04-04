package todos

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"todo-api/utils"

	"github.com/gorilla/mux"
)

func TestCreateTask(t *testing.T) {
	ms := &utils.MockStore{}
	service := NewTodosService(ms)
	t.Run("should return todos on success", func(t *testing.T) {
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
}
