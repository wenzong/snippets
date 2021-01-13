// +build integration

package main

import (
	"context"
	"testing"

	pool "github.com/jolestar/go-commons-pool/v2"
)

func TestPool(t *testing.T) {
	ctx := context.Background()
	p := pool.NewObjectPoolWithDefaultConfig(ctx, &SimpleProducerFactory{url: urlFromEnv()})
	p.Config.LIFO = false
	p.Config.TestOnBorrow = true
	p.Config.MaxTotal = 8

	for i := 0; i < 100; i++ {
		_, err := p.BorrowObject(ctx)
		if err != nil {
			panic(err)
		}
		p.ReturnObject(ctx, obj)
	}
}
