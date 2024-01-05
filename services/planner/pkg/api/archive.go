package api

import "github.com/dsbasko/yandex-go-diploma-1/services/planner/internal/entities"

type ArchiveTaskRequestV1 struct {
	IsArchive bool `json:"is_archive"`
}
type ArchiveTaskResponseV1 entities.RepositoryTaskEntity
