package grpc

import (
	"log"
	"net"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Option struct {
	Listen string
}

func NewListener(v *viper.Viper) net.Listener {
	var o Option
	if err := v.UnmarshalKey("grpc", &o); err != nil {
		panic(errors.Wrap(err, "unmarshal gRPC option failed"))
	}

	lis, err := net.Listen("tcp", o.Listen)
	if err != nil {
		panic(errors.Wrapf(err, "listen on gRPC %s failed ", o.Listen))
	}
	log.Printf("gRPC listening on %s", o.Listen)

	return lis
}
