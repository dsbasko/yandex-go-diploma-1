package task

import (
	"context"
	"fmt"
	"time"

	"github.com/dsbasko/yandex-go-diploma-1/services/planner/pkg/api"
)

func (s *Service) FindToday(ctx context.Context, userID string) (*api.GetTodayResponseV1, error) {
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

	var response []api.GetTodayResponseV1Data
	for _, resp := range *repoResponse {
		response = append(response, api.GetTodayResponseV1Data{
			UUID:        resp.ID,
			UserID:      resp.UserID,
			Name:        resp.Name,
			Description: resp.Description,
			DueDate:     resp.DueDate,
			CreatedAt:   resp.CreatedAt,
			UpdatedAt:   resp.UpdatedAt,
		})
	}

	return &api.GetTodayResponseV1{
		Data:  response,
		Total: len(response),
	}, nil
}
