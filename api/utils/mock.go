package utils

import "todo-api/types"

type MockStore struct{}

func (ms *MockStore) GetTodos() ([]types.Todo, error) {
	return []types.Todo{}, nil
}
