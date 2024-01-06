package api

import (
	"time"

	"github.com/dsbasko/yandex-go-diploma-1/services/planner/internal/entities"
)

type ChangeDueDateRequestV1 struct {
	DueDate time.Time `json:"due_date"`
}
type ChangeDueDateResponseV1 entities.RepositoryTaskEntity

/*
Трансформация структур
*/

func ChangeDueDateResponseV1FromEntity(entity *entities.RepositoryTaskEntity) *ChangeDueDateResponseV1 {
	return &ChangeDueDateResponseV1{
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
