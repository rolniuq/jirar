// Package cli provides the main CLI application structure and command routing.
package cli

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"jirar/internal/config"
)

// App represents the CLI application.
type App struct {
	ctx    context.Context
	logger *logrus.Logger
	config *config.Config
	root   *cobra.Command
}

// NewApp creates a new CLI application instance.
func NewApp(ctx context.Context, logger *logrus.Logger, cfg *config.Config) *App {
	app := &App{
		ctx:    ctx,
		logger: logger,
		config: cfg,
	}

	app.root = app.buildRootCommand()

	return app
}

// Run executes the CLI application.
func (a *App) Run() error {
	return a.root.Execute()
}

// buildRootCommand creates the root Cobra command.
func (a *App) buildRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "jirar",
		Short: "A modern CLI tool for Jira",
		Long: `Jirar is a fast, modern CLI tool for managing Jira tickets 
and receiving notifications directly from your terminal.

Get started with:
  jirar config init    # Interactive setup
  jirar list           # List your tickets`,
		Version:      "0.1.0",
		SilenceUsage: true,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			a.setupLogging()
		},
	}

	// Global flags
	cmd.PersistentFlags().BoolVar(&a.config.Debug, "debug", a.config.Debug, "Enable debug logging")
	cmd.PersistentFlags().StringVar(&a.config.LogLevel, "log-level", a.config.LogLevel, "Log level (debug, info, warn, error)")

	// Add subcommands
	cmd.AddCommand(
		a.buildListCommand(),
		a.buildSearchCommand(),
		a.buildOpenCommand(),
		a.buildConfigCommand(),
	)

	return cmd
}

// setupLogging configures the logger based on configuration.
func (a *App) setupLogging() {
	if a.config.IsDebug() {
		a.logger.SetLevel(logrus.DebugLevel)
	} else {
		switch a.config.GetLogLevel() {
		case "debug":
			a.logger.SetLevel(logrus.DebugLevel)
		case "warn":
			a.logger.SetLevel(logrus.WarnLevel)
		case "error":
			a.logger.SetLevel(logrus.ErrorLevel)
		default:
			a.logger.SetLevel(logrus.InfoLevel)
		}
	}
}
