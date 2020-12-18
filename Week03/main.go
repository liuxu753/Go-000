package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func startHelloWorldService(ctx context.Context) error {
	svr := http.Server{Addr: ":8000"}
	go func() {
		select {
		case <-ctx.Done():
			timeOut, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancel()
			svr.Shutdown(timeOut)
		}
	}()
	err := svr.ListenAndServe()
	return err
}

func startDebugService(ctx context.Context) error {
	svr := http.Server{Addr: ":8001"}
	go func() {
		select {
		case <-ctx.Done():
			timeOut, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancel()
			svr.Shutdown(timeOut)
		}
	}()
	err := svr.ListenAndServe()
	return err
}

func waitOsQuitSignal(ctx context.Context) error {
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-quit:
		return fmt.Errorf("os:kill process")
	case <-ctx.Done():
		return nil
	}
}

func main() {
	group, errCtx := errgroup.WithContext(context.Background())
	group.Go(func() error {
		return waitOsQuitSignal(errCtx)
	})
	group.Go(func() error {
		return startHelloWorldService(errCtx)
	})
	group.Go(func() error {
		return startDebugService(errCtx)
	})

	if err := group.Wait(); err != nil {
		fmt.Printf("error:%v", err)
	}
	return
}
