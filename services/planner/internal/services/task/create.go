package task

import (
	"context"
	"fmt"

	"github.com/dsbasko/yandex-go-diploma-1/services/planner/pkg/api"
)

func (s *Service) Create(ctx context.Context, dto *api.CreateTaskRequestV1) (*api.CreateTaskResponseV1, error) {
	if ctx == nil || dto == nil {
		return nil, ErrArgumentsNotFilled
	}

	switch {
	case dto.UserID == "":
		return nil, ErrEmptyUserID
	case len(dto.Name) < NameMinLength:
		return nil, ErrNameMinLength
	case len(dto.Name) > NameMaxLength:
		return nil, ErrNameMaxLength
	default:
	}

	response, err := s.repo.CreateTask(ctx, dto)
	if err != nil {
		return nil, fmt.Errorf("repo.CreateTask: %w", err)
	}

	return &api.CreateTaskResponseV1{
		UUID: response.ID,
		Name: response.Name,
	}, nil
}
