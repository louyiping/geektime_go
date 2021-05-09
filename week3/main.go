package week3

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"golang.org/x/sync/errgroup"
	"syscall"
	"time"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())

	mux := http.NewServeMux()
	shutdownChan := make(chan struct{})
	mux.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
		shutdownChan <- struct{}{}
	})

	server := http.Server{
		Handler: mux,
		Addr:    ":8080",
	}

	g.Go(func() error {
		return server.ListenAndServe()
	})

	g.Go(func() error {
		select {
		case <-ctx.Done():
			log.Println("errgroup exit")
		case <-shutdownChan:
			log.Println("server will shutdown")
		}

		timeoutCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		log.Println("shutting down server...")
		return server.Shutdown(timeoutCtx)
	})

	g.Go(func() error {
		quit := make(chan os.Signal, 0)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		select {
		case <-ctx.Done():
			return ctx.Err()
		case sig := <-quit:
			return errors.Errorf("get os signal: %v", sig)
		}
	})
	err := g.Wait()
	if err !=nil{
		log.Fatalf("wait errgroup error:%v",err)
	}

}
