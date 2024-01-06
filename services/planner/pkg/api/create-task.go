package api

import (
	"time"

	"github.com/dsbasko/yandex-go-diploma-1/services/planner/internal/entities"
)

type CreateTaskRequestV1 struct {
	UserID      string    `json:"user_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
}

type CreateTaskResponseV1 entities.RepositoryTaskEntity

/*
Трансформация структур
*/

func CreateTaskResponseV1FromEntity(entity *entities.RepositoryTaskEntity) *CreateTaskResponseV1 {
	return &CreateTaskResponseV1{
		ID:          entity.ID,
		UserID:      entity.UserID,
		Name:        entity.Name,
		Description: entity.Description,
		DueDate:     entity.DueDate,
		IsArchive:   entity.IsArchive,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
	}
}
