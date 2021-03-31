package app

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type App struct {
	httpServer *http.Server
	grpcServer *grpc.Server
	listener   net.Listener

	// NOTE: more goroutines maintained here
	//
	// viper.ReadRemoteConfig/WatchRemoteConfig
}

func NewApp(s *http.Server, g *grpc.Server, l net.Listener) *App {
	return &App{
		httpServer: s,
		grpcServer: g,
		listener:   l,
	}
}

func (app *App) Run(ctx context.Context) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	if app.httpServer != nil {
		go func() {
			if err := app.httpServer.ListenAndServe(); err != http.ErrServerClosed {
				panic(errors.Wrap(err, "HTTP server error"))
			}
			log.Println("HTTP Server stopped")
		}()
	}

	if app.grpcServer != nil {
		go func() {
			if err := app.grpcServer.Serve(app.listener); err != nil {
				panic(errors.Wrap(err, "gRPC server error"))
			}
			log.Println("gRPC Server stopped")
		}()
	}

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if app.httpServer != nil {
		if err := app.httpServer.Shutdown(ctx); err != nil {
			log.Fatalf("HTTP Server shutdown error: %+v", err)
		}
		log.Println("HTTP Server exit.")
	}

	if app.grpcServer != nil {
		app.grpcServer.GracefulStop()
		log.Println("gRPC Server exit.")
	}

	log.Println("App stopped")
}

var ProviderSet = wire.NewSet(NewApp)
