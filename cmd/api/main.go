package main

import (
	"context"
	"net/http"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"mini-e-wallet/internal/shared/global"
	"mini-e-wallet/internal/shared/initialize"
)

func main() {
	appCtx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	initialize.Init()

	// engine := initialize.InitRouter()

	port := strconv.Itoa(global.Config.Server.Port)

	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      nil,
		ReadTimeout:  global.Config.Server.ReadTimeout,
		WriteTimeout: global.Config.Server.WriteTimeout,
		IdleTimeout:  global.Config.Server.IdleTimeout,
	}

	go func() {
		global.Logger.Info(context.Background(), "Server is running on port "+port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.Logger.Fatal(context.Background(), "server failed to start: "+err.Error())
		}
	}()

	<-appCtx.Done()

	global.Logger.Info(context.Background(), "Server is shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		global.Logger.Fatal(context.Background(), "server forced to shutdown: "+err.Error())
	}

	initialize.Close(ctx)
	global.Logger.Info(context.Background(), "Server exited properly")
}
