package consumers

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/dsbasko/yandex-go-diploma-1/core/rmq"
	"github.com/dsbasko/yandex-go-diploma-1/services/auth/pkg/api"
	"github.com/rabbitmq/amqp091-go"
)

func reply(
	ctx context.Context,
	conn *rmq.Connector,
	msg *amqp091.Delivery,
	resp *api.JWTValidationResponseV1,
) error {
	respBytes, err := json.Marshal(resp)
	if err != nil {
		return fmt.Errorf("json.Marshal: %w", err)
	}

	if err = conn.Publish(ctx, &rmq.PublisherConfig{
		Key:       msg.ReplyTo,
		Mandatory: true,
		Msg: amqp091.Publishing{
			Body: respBytes,
		},
	}); err != nil {
		return fmt.Errorf("conn.Publish: %w", err)
	}

	return nil
}
