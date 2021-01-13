package main

import "github.com/streadway/amqp"

// One Connection and One Channel
type SimpleProducer struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

func (p *SimpleProducer) Publish(exchange, routingKey string, body []byte) error {
	return p.Channel.Publish(
		exchange,   // exchange
		routingKey, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Transient,
			Body:         body,
		})
}

func NewSimpleProducer(url string) (*SimpleProducer, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &SimpleProducer{Connection: conn, Channel: channel}, nil
}
