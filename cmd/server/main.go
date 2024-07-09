package main

import (
	"context"
	"log/slog"

	"github.com/anurag925/billion/cmd"
	"github.com/anurag925/billion/configs"
	"github.com/millionsmonitoring/millionsgocore/env"
)

func main() {
	ctx := context.Background()
	slog.InfoContext(ctx, "Starting the application from the main...")
	cmd.BasicConfig(ctx)
	slog.InfoContext(ctx, "Application is starting on", "app_name", *configs.CLIFlags().AppName, "environment", env.Env(), "port", *configs.CLIFlags().Port)
	if env.IsDevelopment() {
		slog.InfoContext(ctx, "the settings are", "settings", configs.Settings())
	}
	// routes.Init()

	// // Start server
	// go func() {
	// 	if err := core.Server().Start(fmt.Sprintf(":%d", *flags.Port)); err != nil && err != http.ErrServerClosed {
	// 		slog.Error("shutting down the server", err)
	// 		panic("unable to start the server")
	// 	}
	// }()

	// // Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// // Use a buffered channel to avoid missing signals as recommended for signal.Notify
	// quit := make(chan os.Signal, 1)
	// signal.Notify(quit, os.Interrupt)
	// <-quit
	// ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	// defer cancel()
	// if err := core.Server().Shutdown(ctx); err != nil {
	// 	slog.Error("unable to gracefully stop the server ", "error", err)
	// 	panic("unable to gracefully stop the server")
	// }
}
