package api

import "time"

type CreateTaskRequestV1 struct {
	UserID      string    `json:"user_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
}

type CreateTaskResponseV1 struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}
