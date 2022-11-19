package main

import (
	"errors"
	deliveryHttp "github.com/onemgvv/fakapi/internal/delivery/http"
	"github.com/onemgvv/fakapi/internal/server"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	handler := deliveryHttp.NewHandler()
	app := server.NewServer(handler.InitRoutes())

	go func() {
		if err := app.Run(); errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("[Server]: %s", err.Error())
		}
	}()

	println("Server started")

	/**
	 *	Graceful Shutdown
	 */
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := app.Stop(); err != nil {
		log.Fatalf("Shutdown error")
	}

	log.Println("Server stopped!")
}
