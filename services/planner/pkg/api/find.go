package api

import "time"

type GetTodayResponseV1 struct {
	Data  []GetTodayResponseV1Data `json:"data,omitempty"`
	Total int                      `json:"total"`
}

type GetTodayResponseV1Data struct {
	UUID        string    `json:"uuid"`
	UserID      string    `json:"user_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
