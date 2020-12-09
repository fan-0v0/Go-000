package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

func main() {
	stop := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	g, _ := errgroup.WithContext(ctx)
	server := http.Server{
		Addr: ":8081",
	}
	g.Go(func() error {
		return server.ListenAndServe()
	})

	g.Go(func() error {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT, os.Interrupt)
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-c:
				return errors.New("quit signal")
			}
		}
	})

	err := g.Wait()
	fmt.Println(err)
	<-stop
	fmt.Println("server stopped")
}
