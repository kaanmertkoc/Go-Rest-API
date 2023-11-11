package types

import "time"

const TODO_COLLECTION = "todos"

type Todo struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Completed bool `json:"completed"`
	CreatedAt time.Time`json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}