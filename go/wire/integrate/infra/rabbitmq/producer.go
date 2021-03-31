package rabbitmq

import (
	"context"

	"github.com/pkg/errors"
	"github.com/streadway/amqp"
)

type Producer struct {
	pool *Pool
}

func (p *Producer) Pub(ctx context.Context, exchange, routingKey string, message []byte) error {
	sess, cleanup, err := p.pool.Borrow(ctx)
	if err != nil {
		return err
	}
	defer cleanup(ctx)

	err = sess.Channel.Publish(
		exchange,
		routingKey,
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Transient,
			Body:         message,
		},
	)
	if err != nil {
		return errors.Wrap(err, "Publish message to RabbitMQ failed.")
	}

	return nil
}

func NewProducer(p *Pool) *Producer {
	return &Producer{
		pool: p,
	}
}
