package api

import "github.com/dsbasko/yandex-go-diploma-1/services/planner/internal/entities"

type ChangeIsArchiveRequestV1 struct {
	IsArchive bool `json:"is_archive"`
}
type ChangeIsArchiveResponseV1 entities.RepositoryTaskEntity
