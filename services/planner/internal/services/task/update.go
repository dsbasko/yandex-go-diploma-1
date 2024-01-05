package task

import (
	"context"
	"fmt"

	"github.com/dsbasko/yandex-go-diploma-1/services/planner/pkg/api"
)

func (s *Service) UpdateOnce(ctx context.Context, userID, id string, dto *api.UpdateTaskRequestV1) (*api.UpdateTaskResponseV1, error) {
	if ctx == nil || dto == nil {
		return nil, ErrArgumentsNotFilled
	}

	switch {
	case userID == "":
		return nil, ErrEmptyUserID
	case id == "":
		return nil, ErrEmptyID
	case len(dto.Name) > 0 && len(dto.Name) < NameMinLength:
		return nil, ErrNameMinLength
	case len(dto.Name) > 0 && len(dto.Name) > NameMaxLength:
		return nil, ErrNameMaxLength
	default:
	}

	response, err := s.repo.UpdateOnce(ctx, userID, id, dto)
	if err != nil {
		return nil, fmt.Errorf("repo.UpdateOnce: %w", err)
	}

	return &api.UpdateTaskResponseV1{
		ID:          response.ID,
		UserID:      response.UserID,
		Name:        response.Name,
		Description: response.Description,
		DueDate:     response.DueDate,
		IsArchive:   response.IsArchive,
		CreatedAt:   response.CreatedAt,
		UpdatedAt:   response.UpdatedAt,
	}, nil
}

func (s *Service) UpdateIsArchive(
	ctx context.Context,
	userID, id string,
	dto *api.ChangeIsArchiveRequestV1,
) (*api.ChangeIsArchiveResponseV1, error) {
	if ctx == nil || dto == nil {
		return nil, ErrArgumentsNotFilled
	}

	switch {
	case userID == "":
		return nil, ErrEmptyUserID
	case id == "":
		return nil, ErrEmptyID
	default:
	}

	response, err := s.repo.UpdateIsArchive(ctx, userID, id, dto.IsArchive)
	if err != nil {
		return nil, fmt.Errorf("repo.UpdateOnce: %w", err)
	}

	return &api.ChangeIsArchiveResponseV1{
		ID:          response.ID,
		UserID:      response.UserID,
		Name:        response.Name,
		Description: response.Description,
		DueDate:     response.DueDate,
		IsArchive:   response.IsArchive,
		CreatedAt:   response.CreatedAt,
		UpdatedAt:   response.UpdatedAt,
	}, nil
}

func (s *Service) UpdateDueDate(
	ctx context.Context,
	userID, id string,
	dto *api.ChangeDueDateRequestV1,
) (*api.ChangeDueDateResponseV1, error) {
	if ctx == nil || dto == nil {
		return nil, ErrArgumentsNotFilled
	}

	switch {
	case userID == "":
		return nil, ErrEmptyUserID
	case id == "":
		return nil, ErrEmptyID
	default:
	}

	response, err := s.repo.UpdateDueDate(ctx, userID, id, dto.DueDate)
	if err != nil {
		return nil, fmt.Errorf("repo.UpdateOnce: %w", err)
	}

	return &api.ChangeDueDateResponseV1{
		ID:          response.ID,
		UserID:      response.UserID,
		Name:        response.Name,
		Description: response.Description,
		DueDate:     response.DueDate,
		IsArchive:   response.IsArchive,
		CreatedAt:   response.CreatedAt,
		UpdatedAt:   response.UpdatedAt,
	}, nil
}
