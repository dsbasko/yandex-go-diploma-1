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

	response, err := s.repo.Create(ctx, dto)
	if err != nil {
		return nil, fmt.Errorf("repo.Create: %w", err)
	}

	return &api.CreateTaskResponseV1{
		ID:   response.ID,
		Name: response.Name,
	}, nil
}
