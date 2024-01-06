package api

import "github.com/dsbasko/yandex-go-diploma-1/services/planner/internal/entities"

type GetTasksResponseV1 struct {
	Data  []GetTaskResponseV1 `json:"data"`
	Total int                 `json:"total"`
}

type GetTaskResponseV1 entities.RepositoryTaskEntity

/*
Трансформация структур
*/

func GetTaskResponseV1FromEntity(entity entities.RepositoryTaskEntity) *GetTaskResponseV1 {
	return &GetTaskResponseV1{
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
