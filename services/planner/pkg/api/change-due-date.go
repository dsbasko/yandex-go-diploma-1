package api

import (
	"time"

	"github.com/dsbasko/yandex-go-diploma-1/services/planner/internal/entities"
)

type ChangeDueDateRequestV1 struct {
	DueDate time.Time `json:"due_date"`
}
type ChangeDueDateResponseV1 entities.RepositoryTaskEntity
