package main

import (
	"flag"
	"github.com/avegao/gocondi"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"fmt"
	"github.com/heroku/rollrus"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"context"
	_ "github.com/avegao/iot-api/docs"
	_ "github.com/lib/pq"
)

const (
	version    = "1.0.0"
	apiVersion = "1.0"
)

var (
	debug           = flag.Bool("debug", false, "Print debug logs")
	openevseAddress = flag.String("openevse_address", "openevse:50000", "The server address in the format of host:port")
	buildDate       string
	commitHash      string
	container       *gocondi.Container
	parameters      map[string]interface{}
	server          *http.Server
)

func initContainer() {
	flag.Parse()

	parameters = map[string]interface{}{
		"build_date":       buildDate,
		"debug":            *debug,
		"commit_hash":      commitHash,
		"openevse_address": *openevseAddress,
		"version":          version,
	}

	logger := initLogger()
	gocondi.Initialize(logger)
	container = gocondi.GetContainer()

	for name, value := range parameters {
		container.SetParameter(name, value)
	}
}

func initLogger() *logrus.Logger {
	logLevel := logrus.InfoLevel
	environment := "release"
	log := logrus.New()
	ginMode := gin.ReleaseMode

	if *debug {
		logLevel = logrus.DebugLevel
		environment = "debug"
		ginMode = gin.DebugMode
	} else {
		hook := rollrus.NewHook(fmt.Sprintf("%v", parameters["rollbar_token"]), environment)
		log.Hooks.Add(hook)
	}

	gin.SetMode(ginMode)

	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetFormatter(&logrus.TextFormatter{})
	log.SetLevel(logLevel)

	return log
}

func handleInterrupt() {
	gracefulStop := make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	go func() {
		<-gracefulStop
		powerOff()
	}()
}

func powerOff() {
	container.GetLogger().Infof("Shutting down...")
	closeHttpServer()

	os.Exit(0)
}

func initHttpServer() {
	router := initRouter()
	router.Run(":8080")

	go func() {
		if err := http.ListenAndServe(":8080", router); err != nil {
			container.GetLogger().WithError(err).Panicf("Error creating server")
		} else {
			container.GetLogger().Infof("Listening to 0.0.0.0:8080")
		}
	}()
}

func closeHttpServer() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		container.GetLogger().WithError(err).Fatalf("Server shutdown error")
	}

	container.GetLogger().Debugf("HTTP server closed")
}

// @title IoT API
// @version 1.0.0
// @BasePath /1.0
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	initContainer()
	handleInterrupt()

	logger := container.GetLogger()
	logger.Infof("IoT API v%s started (commit %s, build date %s)", container.GetStringParameter("version"), container.GetStringParameter("commit_hash"), container.GetStringParameter("build_date"))

	initHttpServer()
}
