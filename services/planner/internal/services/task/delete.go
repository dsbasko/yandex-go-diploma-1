package task

import (
	"context"
	"fmt"
)

func (s *Service) DeleteByID(ctx context.Context, userID, id string) error {
	if ctx == nil {
		return ErrArgumentsNotFilled
	}

	switch {
	case userID == "":
		return ErrEmptyUserID
	case id == "":
		return ErrEmptyID
	default:
	}

	if err := s.repo.DeleteByID(ctx, userID, id); err != nil {
		return fmt.Errorf("repo.DeleteByID: %w", err)
	}

	return nil
}
