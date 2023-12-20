package consumers

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/dsbasko/yandex-go-diploma-1/core/logger"
	"github.com/dsbasko/yandex-go-diploma-1/core/rmq"
	"github.com/dsbasko/yandex-go-diploma-1/services/auth/internal/services/jwt"
	"github.com/dsbasko/yandex-go-diploma-1/services/auth/pkg/api"
	"github.com/rabbitmq/amqp091-go"
)

func Validation(ctx context.Context, log *logger.Logger, jwtService *jwt.Service, conn *rmq.Connector) error {
	pullMsgFn := func(msg amqp091.Delivery) *amqp091.Publishing {
		var request api.JWTValidationRequestV1
		var response api.JWTValidationResponseV1
		var replyMsg amqp091.Publishing

		if err := json.Unmarshal(msg.Body, &request); err != nil {
			log.Errorf("json.Unmarshal: %v", err)
			return &replyMsg
		}

		validate, err := jwtService.Validation(ctx, &request)
		if err != nil {
			log.Errorf("jwtService.Validation: %v", err)
			return &replyMsg
		}
		response = *validate

		body, err := json.Marshal(response)
		if err != nil {
			log.Errorf("json.Marshal: %v", err)
			return &replyMsg
		}

		replyMsg.Body = body
		if err = conn.SimplePublishReply(ctx, &rmq.SimplePublisherReplyConfig{
			IncomingMsg: msg,
			ReplyMsg:    replyMsg,
			Mandatory:   true,
		}); err != nil {
			log.Errorf("conn.SimplePublishReply: %v", err)
			return &replyMsg
		}

		return &replyMsg
	}

	err := conn.SimpleConsume(ctx, &rmq.SimpleConsumeConfig{
		Exchange:  api.AMQPExchange,
		Queue:     api.AMQPQueueJWTValidation,
		Key:       api.AMQPKeyJWTValidation,
		Consumer:  "auth.validation.consumer",
		PullMsgFn: pullMsgFn,
	})
	if err != nil {
		return fmt.Errorf("conn.SimpleConsume: %w", err)
	}

	return nil
}
