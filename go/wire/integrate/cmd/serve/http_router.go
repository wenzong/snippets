package serve

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wenzong/demo/biz/user"
	"github.com/wenzong/demo/infra/http/middlewares"
)

// Router assemble all controllers and provider http.Handler for http server
func Router(userController *user.Controller) http.Handler {
	router := mux.NewRouter().StrictSlash(true)
	router.Methods("GET").Path("/user").Name("user").Handler(middlewares.PrivateIPMiddleWare(http.HandlerFunc(userController.Get)))

	return router
}
