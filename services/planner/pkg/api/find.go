package api

import "github.com/dsbasko/yandex-go-diploma-1/services/planner/internal/entities"

type GetTasksResponseV1 struct {
	Data  []GetTaskResponseV1 `json:"data"`
	Total int                 `json:"total"`
}

type GetTaskResponseV1 entities.RepositoryTaskEntity
