package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/google/wire"
	"github.com/pkg/errors"
)

type App struct {
	httpServer *http.Server

	// NOTE: more goroutines maintained here
	//
	// grpcServer
	// viper.ReadRemoteConfig/WatchRemoteConfig
	// amqp auto reconnect
}

func NewApp(s *http.Server) *App {
	return &App{
		httpServer: s,
	}
}

func (app *App) Run(ctx context.Context) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := app.httpServer.ListenAndServe(); err != http.ErrServerClosed {
			panic(errors.Wrap(err, "HTTP server error"))
		}
		log.Println("HTTP Server stopped")
	}()

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := app.httpServer.Shutdown(ctx); err != nil {
		log.Fatalf("HTTP Server shutdown error: %+v", err)
	}

	log.Println("App stopped")
}

var ProviderSet = wire.NewSet(NewApp)
