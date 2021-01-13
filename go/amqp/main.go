package main

import (
	"context"
	"os"

	pool "github.com/jolestar/go-commons-pool/v2"
	"github.com/streadway/amqp"
)

func urlFromEnv() string {
	url := os.Getenv("AMQP_URL")
	if url == "" {
		url = "amqp://"
	}
	return url
}

func main() {
	ctx := context.Background()
	p := pool.NewObjectPoolWithDefaultConfig(ctx, &ConnectionFactory{url: urlFromEnv()})
	p.Config.LIFO = false
	p.Config.TestOnBorrow = true
	p.Config.TestOnReturn = true

	// make some connections
	for i := 0; i < 8; i++ {
		go func() {
			obj, err := p.BorrowObject(ctx)
			if err != nil {
				panic(err)
			}

			p.ReturnObject(ctx, obj)
		}()
	}

	for i := 0; i < 100; i++ {
		// Connection Pool
		conn, err := p.BorrowObject(ctx)
		if err != nil {
			panic(err)
		}
		p.ReturnObject(ctx, conn)

		// Channel Pool
		channel, err := conn.(*ConnectionWithChannelPool).ChannelPool.BorrowObject(ctx)
		if err != nil {
			panic(err)
		}
		conn.(*ConnectionWithChannelPool).ChannelPool.ReturnObject(ctx, channel)

		err = channel.(*SimpleProducer).Channel.Publish(
			"amq.topic", // exchange
			"event",     // routing key
			false,       // mandatory
			false,       // immediate
			amqp.Publishing{
				DeliveryMode: amqp.Transient,
				Body:         []byte("hello"),
			})
		if err != nil {
			panic(err)
		}
	}
}
