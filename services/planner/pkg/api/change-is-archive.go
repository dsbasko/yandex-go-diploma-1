package api

import "github.com/dsbasko/yandex-go-diploma-1/services/planner/internal/entities"

type ChangeIsArchiveRequestV1 struct {
	IsArchive bool `json:"is_archive"`
}
type ChangeIsArchiveResponseV1 entities.RepositoryTaskEntity

/*
Трансформация структур
*/

func ChangeIsArchiveResponseV1FromEntity(entity *entities.RepositoryTaskEntity) *ChangeIsArchiveResponseV1 {
	return &ChangeIsArchiveResponseV1{
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
