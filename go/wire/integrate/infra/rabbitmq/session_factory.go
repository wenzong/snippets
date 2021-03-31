package rabbitmq

import (
	"context"

	pool "github.com/jolestar/go-commons-pool/v2"
)

// SessionFactory new Session from dsn, and wrap it to pool.PooledObject
type SessionFactory struct {
	dsn string
}

func (sf *SessionFactory) MakeObject(ctx context.Context) (*pool.PooledObject, error) {
	sess, err := NewSession(sf.dsn)
	return pool.NewPooledObject(sess), err
}

func (sf *SessionFactory) DestroyObject(ctx context.Context, object *pool.PooledObject) error {
	object.Object.(*Session).Connection.Close()
	return nil
}

func (sf *SessionFactory) ValidateObject(ctx context.Context, object *pool.PooledObject) bool {
	return !object.Object.(*Session).Connection.IsClosed()
}

func (sf *SessionFactory) ActivateObject(ctx context.Context, object *pool.PooledObject) error {
	return nil
}

func (sf *SessionFactory) PassivateObject(ctx context.Context, object *pool.PooledObject) error {
	return nil
}

func NewSessionFactory(dsn string) *SessionFactory {
	return &SessionFactory{dsn: dsn}
}
