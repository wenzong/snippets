package main

import (
	"time"

	"github.com/pkg/errors"
	"github.com/streadway/amqp"
)

// Auto reconnect on connection close
type ReconnectProducer struct {
	url    string
	logger Logger

	errChan    chan *amqp.Error
	connection *amqp.Connection
	channel    *amqp.Channel

	closed bool
}

func NewReconnectProducer(url string, logger Logger) *ReconnectProducer {
	p := &ReconnectProducer{
		url:    url,
		logger: logger,
	}
	p.connect()
	go p.reconnector()

	return p
}

func (p *ReconnectProducer) Pub(exchange, routingKey string, body []byte) error {
	if p.closed {
		return errors.New("AMQP Producer: already closed")
	}

	err := p.channel.Publish(
		exchange,   // exchange
		routingKey, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Transient,
			Body:         body,
		})
	if err != nil {
		p.logger.Warnf("AMQP Producer: publish message failed: %v", err)
	}

	return err
}

func (p *ReconnectProducer) Close() {
	p.logger.Info("AMQP Producer: closing.")
	p.closed = true
	p.channel.Close()
	p.connection.Close()
}

func (p *ReconnectProducer) reconnector() {
	for {
		err := <-p.errChan
		if !p.closed {
			p.logger.Infof("AMQP Producer: unexpected connection close, reconnecting: %v", err)
			p.connect()
		} else {
			p.logger.Info("AMQP Producer: reconnector goroutine stopping.")
			return
		}
	}
}

func (p *ReconnectProducer) connect() {
	for {
		if p.closed {
			return
		}

		p.logger.Info("AMQP Producer: connecting")
		conn, err := amqp.Dial(p.url)
		if err == nil {
			p.connection = conn
			channel, err := p.connection.Channel()
			if err == nil {
				p.channel = channel

				p.errChan = make(chan *amqp.Error)
				p.connection.NotifyClose(p.errChan)
				return
			}

			p.logger.Warnf("AMQP Producer: open channel failed: %v", err)
			p.connection.Close()
		}

		p.logger.Warnf("AMQP Producer: connect failed, retrying in 1 sec: %v", err)
		time.Sleep(1000 * time.Millisecond)
	}
}
