package rabbitmq

import (
	"github.com/streadway/amqp"
)

type Session struct {
	*amqp.Connection
	*amqp.Channel
}

func NewSession(url string) (*Session, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	go func() {
		<-channel.NotifyClose(make(chan *amqp.Error))
		conn.Close()
	}()

	return &Session{Connection: conn, Channel: channel}, nil
}
