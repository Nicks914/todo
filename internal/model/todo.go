package model

import "time"

type Todo struct {
	ID        string    `json:"id"`
	Task      string    `json:"task"`
	DueDate   time.Time `json:"due_date"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}
