package amqp

import (
	"fmt"

	"github.com/dsbasko/yandex-go-diploma-1/core/rmq"
	"github.com/dsbasko/yandex-go-diploma-1/services/planner/internal/config"
)

type Adapter struct {
	conn *rmq.Connector
}

func RunAdapter() (*Adapter, error) {
	conn, err := rmq.Connect(config.GetRmqConnectingString())
	if err != nil {
		return nil, fmt.Errorf("rmq.Connect: %w", err)
	}

	return &Adapter{conn: conn}, nil
}
