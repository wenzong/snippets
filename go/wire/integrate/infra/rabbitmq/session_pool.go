package rabbitmq

import (
	"context"

	pool "github.com/jolestar/go-commons-pool/v2"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Pool struct {
	pool *pool.ObjectPool
}

func (p *Pool) Borrow(ctx context.Context) (*Session, func(context.Context), error) {
	sess, err := p.pool.BorrowObject(ctx)
	if err != nil {
		return nil, nil, errors.Wrap(err, "Borrow RabbitMQ Session from pool failed.")
	}

	return sess.(*Session), func(ctx context.Context) {
		p.pool.ReturnObject(ctx, sess)
	}, nil

}

func NewPool(v *viper.Viper) *Pool {
	c := pool.NewDefaultPoolConfig()
	c.LIFO = false
	c.TestOnBorrow = true
	c.TestOnReturn = true
	c.MaxTotal = 100
	return &Pool{
		pool: pool.NewObjectPool(context.Background(), &SessionFactory{dsn: v.GetString("rabbitmq.dsn")}, c),
	}
}
