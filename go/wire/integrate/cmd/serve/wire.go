// +build wireinject

package serve

import (
	"github.com/google/wire"

	"github.com/wenzong/demo/biz/user"
	"github.com/wenzong/demo/infra/app"
	"github.com/wenzong/demo/infra/config"
	"github.com/wenzong/demo/infra/db"
	"github.com/wenzong/demo/infra/grpc"
	"github.com/wenzong/demo/infra/http"
	"github.com/wenzong/demo/infra/log"
)

var ProviderSet = wire.NewSet(
	Router,
	gRPCServerOptions,
	gRPCRegisterServiceFn,
)

var Set = wire.NewSet(
	app.ProviderSet,
	config.ProviderSet,
	db.ProviderSet,
	grpc.ProviderSet,
	http.ProviderSet,
	log.ProviderSet,
	// biz
	user.ProviderSet,
	ProviderSet,
)

func App() (*app.App, func(), error) {
	panic(wire.Build(Set))
}
