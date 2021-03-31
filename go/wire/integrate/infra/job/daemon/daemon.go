package daemon

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/wenzong/demo/infra/job"
	"github.com/wenzong/demo/infra/log"
)

type Daemon struct {
	logger log.Logger
	task   job.Task
	ticker *time.Ticker

	interceptors []job.TaskInterceptor

	http *http.Server
}

func (job *Daemon) Run(ctx context.Context) error {
	job.logger.Info("Daemon Job start")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := job.http.ListenAndServe(); err != http.ErrServerClosed {
			job.logger.Warnf("HTTP server ListenAndServe: %v", err)
		}
	}()

	handler := job.task
	for _, i := range job.interceptors {
		handler = i(handler)
	}

LOOP:
	for {
		select {
		case <-job.ticker.C:
			handler(ctx)
		case <-quit:
			break LOOP
		}
	}

	newCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	if err := job.http.Shutdown(newCtx); err != nil {
		job.logger.Warnf("HTTP server Shutdown: %v", err)
	}

	job.logger.Info("Daemon Job exit")
	return nil
}

func NewDaemon(
	l log.Logger,
	t job.Task,
	ticker *time.Ticker,
	is []job.TaskInterceptor,
	http *http.Server,
) job.Job {
	return &Daemon{
		logger:       l,
		task:         t,
		ticker:       ticker,
		interceptors: is,
		http:         http,
	}
}
