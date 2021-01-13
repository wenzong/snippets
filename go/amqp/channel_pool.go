package main

import (
	"context"

	pool "github.com/jolestar/go-commons-pool/v2"
	"github.com/streadway/amqp"
)

type ConnectionWithChannelPool struct {
	*amqp.Connection
	ChannelPool *pool.ObjectPool
}

// ConnectionFactory create *amqp.Connection(ConnectionWithChannelPool) from url
type ConnectionFactory struct {
	url string
}

func (f *ConnectionFactory) MakeObject(ctx context.Context) (*pool.PooledObject, error) {
	conn, err := amqp.Dial(f.url)
	if err != nil {
		return nil, err
	}

	p := pool.NewObjectPoolWithDefaultConfig(ctx, &ChannelFactory{conn})
	p.Config.TestOnBorrow = true
	return pool.NewPooledObject(&ConnectionWithChannelPool{conn, p}), err
}

func (f *ConnectionFactory) DestroyObject(ctx context.Context, object *pool.PooledObject) error {
	object.Object.(*ConnectionWithChannelPool).Connection.Close()
	object.Object.(*ConnectionWithChannelPool).ChannelPool.Close(ctx)
	return nil
}

func (f *ConnectionFactory) ValidateObject(ctx context.Context, object *pool.PooledObject) bool {
	return !object.Object.(*ConnectionWithChannelPool).IsClosed()
}

func (f *ConnectionFactory) ActivateObject(ctx context.Context, object *pool.PooledObject) error {
	return nil
}

func (f *ConnectionFactory) PassivateObject(ctx context.Context, object *pool.PooledObject) error {
	return nil
}

// ChannelFactory create *amqp.Channel(SimpleProducer) from *amqp.Connection
type ChannelFactory struct {
	conn *amqp.Connection
}

func (f *ChannelFactory) MakeObject(ctx context.Context) (*pool.PooledObject, error) {
	channel, err := f.conn.Channel()
	if err != nil {
		return nil, err
	}
	return pool.NewPooledObject(&SimpleProducer{f.conn, channel}), nil
}

func (f *ChannelFactory) DestroyObject(ctx context.Context, object *pool.PooledObject) error {
	object.Object.(*SimpleProducer).Channel.Close()
	return nil
}

func (f *ChannelFactory) ValidateObject(ctx context.Context, object *pool.PooledObject) bool {
	return !object.Object.(*SimpleProducer).Connection.IsClosed()
}

func (f *ChannelFactory) ActivateObject(ctx context.Context, object *pool.PooledObject) error {
	return nil
}

func (f *ChannelFactory) PassivateObject(ctx context.Context, object *pool.PooledObject) error {
	return nil
}
