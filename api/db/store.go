package db

import (
	"database/sql"
	"todo-api/types"
)

type Store interface {
	GetTodos() ([]types.Todo, error)
	GetTodoById(id string) (types.Todo, error)
}

type Storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{db: db}
}

func (storage *Storage) GetTodos() ([]types.Todo, error) {
	rows, err := storage.db.Query("SELECT * FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todos := []types.Todo{}
	for rows.Next() {
		var todo types.Todo
		err := rows.Scan(&todo.Id, &todo.Name, &todo.Status, &todo.CreatedAt)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func (storage *Storage) GetTodoById(id string) (types.Todo, error) {
	var todo types.Todo
	err := storage.db.QueryRow("SELECT * FROM todos WHERE id = ?", id).Scan(&todo.Id, &todo.Name, &todo.Status, &todo.CreatedAt)
	if err != nil {
		return types.Todo{}, err
	}

	return todo, nil
}
