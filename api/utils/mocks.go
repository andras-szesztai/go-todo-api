package utils

import (
	"errors"
	"todo-api/types"
)

type MockStore struct{}

func (ms *MockStore) GetTodos() ([]types.Todo, error) {
	return []types.Todo{}, nil
}

func (ms *MockStore) GetTodoById(id string) (types.Todo, error) {
	return types.Todo{}, nil
}

func (ms *MockStore) PostTodo(
	todo *types.PostTodoRequest,
) (*types.Todo, error) {
	return &types.Todo{}, nil
}

type MockErrorStore struct{}

func (ms *MockErrorStore) GetTodos() ([]types.Todo, error) {
	return nil, errors.New("error")
}

func (ms *MockErrorStore) GetTodoById(id string) (types.Todo, error) {
	return types.Todo{}, errors.New("error")
}

func (ms *MockErrorStore) PostTodo(
	todo *types.PostTodoRequest,
) (*types.Todo, error) {
	return &types.Todo{}, errors.New("error")
}
