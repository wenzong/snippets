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
)

var ProviderSet = wire.NewSet(
	Router,
	gRPCServerOptions,
	gRPCRegisterServiceFn,
)

var Set = wire.NewSet(
	config.ProviderSet,
	app.ProviderSet,
	db.ProviderSet,
	user.ProviderSet,
	http.ProviderSet,
	grpc.ProviderSet,
	ProviderSet,
)

func App() *app.App {
	panic(wire.Build(Set))
}
