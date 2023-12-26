package rmq

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/rabbitmq/amqp091-go"
)

/*
Отправка сообщения
*/

type PublisherConfig struct {
	Exchange  string
	Key       string
	Mandatory bool
	Msg       amqp091.Publishing
}

func (c *Connector) Publish(
	ctx context.Context,
	cfgPublisher *PublisherConfig,
) error {
	ch, err := c.conn.Channel()
	if err != nil {
		return fmt.Errorf("conn.Channel: %w", err)
	}
	defer ch.Close()

	if cfgPublisher.Msg.CorrelationId == "" {
		cfgPublisher.Msg.CorrelationId = uuid.New().String()
	}

	return ch.PublishWithContext(
		ctx,
		cfgPublisher.Exchange,
		cfgPublisher.Key,
		cfgPublisher.Mandatory,
		false,
		cfgPublisher.Msg,
	)
}

/*
Отправка сообщения и ожидание ответа
*/

type SimplePublisherConfig struct {
	Exchange  string
	Key       string
	Mandatory bool
	Msg       amqp091.Publishing
}

func (c *Connector) SimplePublishAndWaitResponse(
	ctx context.Context,
	cfgPublisher *SimplePublisherConfig,
) ([]byte, error) {
	replyQueue, err := c.QueueDeclare(&QueueConfig{
		Name:       "",
		Durable:    false,
		AutoDelete: true,
		Exclusive:  true,
		NoWait:     false,
		Args:       nil,
	})
	if err != nil {
		return []byte(""), err
	}

	msgs, closeFn, err := c.Consume(ctx, &ConsumeConfig{
		Queue:     replyQueue.Name,
		Consumer:  "",
		AutoAck:   true,
		Exclusive: false,
		NoLocal:   false,
		NoWait:    false,
		Args:      nil,
	})
	defer closeFn()
	if err != nil {
		return []byte(""), err
	}

	if cfgPublisher.Msg.CorrelationId == "" {
		cfgPublisher.Msg.CorrelationId = uuid.New().String()
	}
	if cfgPublisher.Msg.ReplyTo == "" {
		cfgPublisher.Msg.ReplyTo = replyQueue.Name
	}
	if errPub := c.Publish(ctx, &PublisherConfig{
		cfgPublisher.Exchange,
		cfgPublisher.Key,
		cfgPublisher.Mandatory,
		cfgPublisher.Msg,
	}); errPub != nil {
		return []byte(""), errPub
	}

	c.log.Debugw(
		"message has been sent and wait response",
		"payload", string(cfgPublisher.Msg.Body),
		"correlation_id", cfgPublisher.Msg.CorrelationId,
	)

	msg := <-msgs
	c.log.Debugw(
		"a response to the message has been received",
		"payload", string(msg.Body),
		"correlation_id", msg.CorrelationId,
	)

	return msg.Body, nil
}

/*
Отправка сообщения в ответ
*/

type SimplePublisherReplyConfig struct {
	Mandatory   bool
	IncomingMsg amqp091.Delivery
	ReplyMsg    amqp091.Publishing
}

func (c *Connector) SimplePublishReply(
	ctx context.Context,
	cfgPublisher *SimplePublisherReplyConfig,
) error {
	ch, err := c.conn.Channel()
	if err != nil {
		return fmt.Errorf("conn.Channel: %w", err)
	}
	defer ch.Close()

	if cfgPublisher.IncomingMsg.CorrelationId != "" && cfgPublisher.ReplyMsg.CorrelationId == "" {
		cfgPublisher.ReplyMsg.CorrelationId = cfgPublisher.IncomingMsg.CorrelationId
	}

	if err = ch.PublishWithContext(
		ctx,
		"",
		cfgPublisher.IncomingMsg.ReplyTo,
		cfgPublisher.Mandatory,
		false,
		cfgPublisher.ReplyMsg,
	); err != nil {
		return fmt.Errorf("ch.PublishWithContext: %w", err)
	}

	if err = cfgPublisher.IncomingMsg.Ack(true); err != nil {
		return fmt.Errorf("IncomingMsg.Ack: %w", err)
	}

	return nil
}
