package redis

import (
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type DefaultRedis struct {
	*redis.Client
}

type SecondRedis struct {
	*redis.Client
}

func NewDefault(v *viper.Viper) *DefaultRedis {
	var option redis.Options
	config := v.GetStringMap("redis.default")
	err := mapstructure.Decode(config, &option)

	if err != nil {
		panic(errors.Wrap(err, "Parse redis config failed"))
	}

	return &DefaultRedis{redis.NewClient(&option)}
}

func NewSecond(v *viper.Viper) *SecondRedis {
	var option redis.Options
	config := v.GetStringMap("redis.second")
	err := mapstructure.Decode(config, &option)
	if err != nil {
		panic(errors.Wrap(err, "Parse redis config failed"))
	}

	return &SecondRedis{redis.NewClient(&option)}
}

var ProviderSet = wire.NewSet(NewDefault, NewSecond)
