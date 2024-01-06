package interfaces

import (
	"context"
)

//go:generate ../../../../bin/mockgen -destination=../repositories/mock/mock.go -package=mock github.com/dsbasko/yandex-go-diploma-1/services/notification/internal/interfaces Repository

type Repository interface {
	Ping(ctx context.Context) error
}
