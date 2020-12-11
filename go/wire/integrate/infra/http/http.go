package http

import (
	"log"
	"net/http"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Option struct {
	Listen string
}

func NewOption(v *viper.Viper) (o *Option) {
	if err := v.UnmarshalKey("http", &o); err != nil {
		panic(errors.Wrap(err, "Unmarshal HTTP option"))
	}

	return o
}

func NewServer(o *Option, h http.Handler) *http.Server {
	log.Printf("HTTP listening on %s", o.Listen)
	return &http.Server{
		Addr:    o.Listen,
		Handler: h,
	}
}

var ProviderSet = wire.NewSet(NewOption, NewServer)
