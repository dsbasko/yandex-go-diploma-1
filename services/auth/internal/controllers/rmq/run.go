package rmq

import (
	"context"
	"fmt"
	"sync"

	"github.com/dsbasko/yandex-go-diploma-1/core/logger"
	"github.com/dsbasko/yandex-go-diploma-1/core/rmq"
	"github.com/dsbasko/yandex-go-diploma-1/services/auth/internal/config"
	"github.com/dsbasko/yandex-go-diploma-1/services/auth/internal/controllers/rmq/consumers"
	"github.com/dsbasko/yandex-go-diploma-1/services/auth/internal/services/jwt"
	"github.com/dsbasko/yandex-go-diploma-1/services/auth/pkg/api"
)

func RunController(
	ctx context.Context,
	log *logger.Logger,
	jwtService *jwt.Service,
) (func(), error) {
	mu := sync.Mutex{}

	conn, err := rmq.Connect(ctx, log, config.GetRmqConnectingString(), rmq.ConnectorOptions{
		MaxRetries:   config.GetRmqMaxRetries(),
		RetryTimeOut: config.GetRmqRetryTimeout(),
	})
	if err != nil {
		return func() {}, err
	}

	connClose := func() {
		if errClose := conn.Close(); errClose != nil {
			log.Errorf("conn.Close: %v", errClose)
		}
	}

	if err = conn.ExchangeDeclare(&rmq.ExchangeConfig{
		Name:    api.AMQPExchange,
		Kind:    "direct",
		Durable: true,
	}); err != nil {
		return connClose, fmt.Errorf("conn.ExchangeDeclare: %w", err)
	}

	err = consumers.Validation(ctx, &mu, log, jwtService, conn)
	if err != nil {
		return connClose, fmt.Errorf("consumers.Validation: %w", err)
	}

	return connClose, err
}
