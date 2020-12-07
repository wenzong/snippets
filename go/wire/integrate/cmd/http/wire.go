// +build wireinject

package http

import (
	"github.com/google/wire"

	"github.com/wenzong/demo/biz/user"
	"github.com/wenzong/demo/infra/app"
	"github.com/wenzong/demo/infra/config"
	"github.com/wenzong/demo/infra/db"
	"github.com/wenzong/demo/infra/http"
)

var Set = wire.NewSet(
	config.ProviderSet,
	app.ProviderSet,
	db.ProviderSet,
	user.ProviderSet,
	http.ProviderSet,
	Router,
)

func App() *app.App {
	panic(wire.Build(Set))
}
