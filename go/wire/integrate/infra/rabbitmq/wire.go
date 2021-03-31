package rabbitmq

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewProducer, NewPool)
