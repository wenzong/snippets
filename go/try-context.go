package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(ctx context.Context) {
	run := true
	for run {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("sleeping")
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			run = false
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	done := make(chan bool)

	go func() {
		defer close(done)
		Run(ctx)
	}()

	wait := make(chan os.Signal, 1)
	signal.Notify(wait, syscall.SIGINT, syscall.SIGTERM)
	<-wait

	fmt.Println("cancelling")

	// cancel context
	cancel()

	fmt.Println("cancelled")

	<-done
}
