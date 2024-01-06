package api

import "github.com/dsbasko/yandex-go-diploma-1/services/planner/internal/entities"

type UpdateTaskRequestV1 entities.RepositoryTaskEntity

type UpdateTaskResponseV1 entities.RepositoryTaskEntity

/*
Трансформация структур
*/

func UpdateTaskResponseV1FromEntity(entity *entities.RepositoryTaskEntity) *UpdateTaskResponseV1 {
	return &UpdateTaskResponseV1{
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
