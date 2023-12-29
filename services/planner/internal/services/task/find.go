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

	dateStart := time.Now().Truncate(24 * time.Hour) //nolint:gomnd
	dateEnd := dateStart.Add(24*time.Hour - time.Nanosecond)

	repoResponse, err := s.repo.FindByUserIDAndDate(ctx, userID, &dateStart, &dateEnd)
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

func (s *Service) FindWeek(ctx context.Context, userID string) (*api.GetTasksResponseV1, error) {
	if ctx == nil {
		return nil, ErrArgumentsNotFilled
	}

	switch {
	case userID == "":
		return nil, ErrEmptyUserID
	default:
	}

	now := time.Now()
	weekday := int(now.Weekday())
	startOfWeek := now.AddDate(0, 0, -weekday+1).Truncate(24 * time.Hour)         //nolint:gomnd
	endOfWeek := startOfWeek.AddDate(0, 0, 6).Add(24*time.Hour - time.Nanosecond) //nolint:gomnd

	repoResponse, err := s.repo.FindByUserIDAndDate(ctx, userID, &startOfWeek, &endOfWeek)
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

func (s *Service) FindUndated(ctx context.Context, userID string) (*api.GetTasksResponseV1, error) {
	if ctx == nil {
		return nil, ErrArgumentsNotFilled
	}

	switch {
	case userID == "":
		return nil, ErrEmptyUserID
	default:
	}

	repoResponse, err := s.repo.FindByUserIDAndDate(ctx, userID, nil, nil)
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
