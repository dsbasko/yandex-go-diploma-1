package entities

import (
	"time"
)

type RepositoryNotificationEntity struct {
	ID             string    `json:"id"`
	UserID         string    `json:"user_id"`
	TaskID         string    `json:"task_id"`
	NotificationAt time.Time `json:"notification_at"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
