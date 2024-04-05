package types

import "time"

type ErrorResponse struct {
	Error string `json:"error"`
}

type Todo struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
}

type PostTodoRequest struct {
	Name string `json:"name"`
}
