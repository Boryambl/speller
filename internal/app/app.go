package app

import (
	"fmt"
	"homework/pkg/httpserver"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Run() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("logger initialization error")
	}

	// HTTP Server
	handler := gin.New()

	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	r := &routes{logger}
	handler.POST("/api/correct_mistakes", r.correctMistakes)

	httpServer := httpserver.New(handler, httpserver.Port("8080"))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		logger.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		logger.Error("app - Run - httpServer.Notify: " +err.Error())
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		logger.Error(fmt.Sprintf("app - Run - httpServer.Shutdown: %v", err))
	}
}