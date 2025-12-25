// Package main is the entry point for the Jira CLI tool.
package main

import (
	"context"

	"github.com/sirupsen/logrus"

	"jirar/internal/cli"
	"jirar/internal/config"
)

func main() {
	// Initialize logger
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	// Create application context
	ctx := context.Background()

	// Initialize configuration
	cfg, err := config.New()
	if err != nil {
		logger.WithError(err).Fatal("Failed to initialize configuration")
	}

	// Create and run CLI application
	app := cli.NewApp(ctx, logger, cfg)
	if err := app.Run(); err != nil {
		logger.WithError(err).Fatal("Application failed")
	}
}
