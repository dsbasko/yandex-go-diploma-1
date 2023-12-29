package task

import (
	"context"
	"fmt"
	"time"

	"github.com/dsbasko/yandex-go-diploma-1/services/planner/pkg/api"
)

func (s *Service) FindByID(
	ctx context.Context,
	userID, id string,
) (*api.GetTaskResponseV1, error) {
	if ctx == nil {
		return nil, ErrArgumentsNotFilled
	}

	switch {
	case userID == "":
		return nil, ErrEmptyUserID
	case id == "":
		return nil, ErrEmptyID
	default:
	}

	resp, err := s.repo.FindByID(ctx, userID, id)
	if err != nil {
		return nil, fmt.Errorf("repo.FindByID: %w", err)
	}

	if resp == nil {
		return nil, nil
	}

	return &api.GetTaskResponseV1{
		ID:          resp.ID,
		UserID:      resp.UserID,
		Name:        resp.Name,
		Description: resp.Description,
		DueDate:     resp.DueDate,
		IsArchive:   resp.IsArchive,
		CreatedAt:   resp.CreatedAt,
		UpdatedAt:   resp.UpdatedAt,
	}, nil
}

func (s *Service) FindToday(ctx context.Context, userID string) (*api.GetTasksResponseV1, error) {
	if ctx == nil {
		return nil, ErrArgumentsNotFilled
	}

	switch {
	case userID == "":
		return nil, ErrEmptyUserID
	default:
	}

	now := time.Now()
	dateStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	dateEnd := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 1e9-1, now.Location())

	repoResponse, err := s.repo.FindByUserIDAndDate(ctx, userID, dateStart, dateEnd)
	if err != nil {
		return nil, fmt.Errorf("repo.CreateTask: %w", err)
	}

	if repoResponse == nil {
		return &api.GetTasksResponseV1{
			Data:  nil,
			Total: 0,
		}, nil
	}

	var response []api.GetTaskResponseV1
	for _, resp := range *repoResponse {
		response = append(response, api.GetTaskResponseV1{
			ID:          resp.ID,
			UserID:      resp.UserID,
			Name:        resp.Name,
			Description: resp.Description,
			DueDate:     resp.DueDate,
			IsArchive:   resp.IsArchive,
			CreatedAt:   resp.CreatedAt,
			UpdatedAt:   resp.UpdatedAt,
		})
	}

	return &api.GetTasksResponseV1{
		Data:  response,
		Total: len(response),
	}, nil
}

func (s *Service) FindArchive(ctx context.Context, userID string) (*api.GetTasksResponseV1, error) {
	if ctx == nil {
		return nil, ErrArgumentsNotFilled
	}

	switch {
	case userID == "":
		return nil, ErrEmptyUserID
	default:
	}

	repoResponse, err := s.repo.FindArchive(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("repo.CreateTask: %w", err)
	}

	if repoResponse == nil {
		return &api.GetTasksResponseV1{
			Data:  nil,
			Total: 0,
		}, nil
	}

	var response []api.GetTaskResponseV1
	for _, resp := range *repoResponse {
		response = append(response, api.GetTaskResponseV1{
			ID:          resp.ID,
			UserID:      resp.UserID,
			Name:        resp.Name,
			Description: resp.Description,
			DueDate:     resp.DueDate,
			IsArchive:   resp.IsArchive,
			CreatedAt:   resp.CreatedAt,
			UpdatedAt:   resp.UpdatedAt,
		})
	}

	return &api.GetTasksResponseV1{
		Data:  response,
		Total: len(response),
	}, nil
}
