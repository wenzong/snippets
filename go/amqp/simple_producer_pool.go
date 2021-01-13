package main

import (
	"context"

	pool "github.com/jolestar/go-commons-pool/v2"
)

type SimpleProducerFactory struct {
	url string
}

func (f *SimpleProducerFactory) MakeObject(ctx context.Context) (*pool.PooledObject, error) {
	p, err := NewSimpleProducer(f.url)
	return pool.NewPooledObject(p), err
}

func (f *SimpleProducerFactory) DestroyObject(ctx context.Context, object *pool.PooledObject) error {
	object.Object.(*SimpleProducer).Connection.Close()
	return nil
}

func (f *SimpleProducerFactory) ValidateObject(ctx context.Context, object *pool.PooledObject) bool {
	return !object.Object.(*SimpleProducer).Connection.IsClosed()
}

func (f *SimpleProducerFactory) ActivateObject(ctx context.Context, object *pool.PooledObject) error {
	return nil
}

func (f *SimpleProducerFactory) PassivateObject(ctx context.Context, object *pool.PooledObject) error {
	return nil
}
