package app

import (
	_ "embed"
	"fmt"
	json "github.com/json-iterator/go"
	"homework/pkg/httpserver"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//go:embed config.json
var config []byte

func Run() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("logger initialization error")
	}
	port, err := getPort()
	if err != nil {
		logger.Fatal("app - Run - getPort: " + err.Error())
	}

	// HTTP Server
	handler := gin.New()

	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	r := &routes{logger}
	handler.PUT("/api/correct_mistakes", r.correctMistakes)

	httpServer := httpserver.New(handler, httpserver.Port(port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		logger.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		logger.Error("app - Run - httpServer.Notify: " + err.Error())
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		logger.Error(fmt.Sprintf("app - Run - httpServer.Shutdown: %v", err))
	}
}

func getPort() (string, error) {
	var conf struct {
		Port string `json:"port"`
	}
	err := json.Unmarshal(config, &conf)
	return conf.Port, err
}
