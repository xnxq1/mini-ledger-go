package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	httpServer "github.com/xnxq1/mini-ledger-go/internal/http-server"
	"github.com/xnxq1/mini-ledger-go/internal/http-server/merchants"
)

func main() {
	root := chi.NewRouter()
	root.Use(httpServer.ErrorMiddleware)
	initMerchantRouter := merchants.MerchantRouter{}
	root.Mount("/", initMerchantRouter.Init())
	server := &http.Server{
		Addr:         ":8080",
		Handler:      root,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 20 * time.Second,
		IdleTimeout:  50 * time.Second,
	}

	go func() {
		log.Println("Listening on " + server.Addr)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
