package cli

import (
	"github.com/spf13/cobra"
)

// buildListCommand creates the list command.
func (a *App) buildListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List Jira tickets assigned to you",
		Long: `List all Jira tickets that are currently assigned to you.
Supports filtering by status, project, and sorting options.`,
		Run: func(cmd *cobra.Command, args []string) {
			// TODO: Implement list command
			a.logger.Info("List command - coming soon!")
		},
	}

	// Add command flags
	cmd.Flags().StringP("status", "s", "", "Filter by status (todo, in-progress, done)")
	cmd.Flags().IntP("limit", "l", 20, "Maximum number of tickets to show")
	cmd.Flags().StringP("project", "p", "", "Filter by project key")
	cmd.Flags().String("sort", "updated", "Sort field (updated, created, priority)")
	cmd.Flags().Bool("json", false, "Output in JSON format")

	return cmd
}

// buildSearchCommand creates the search command.
func (a *App) buildSearchCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "search [jql]",
		Short: "Search Jira tickets using JQL",
		Long: `Search for Jira tickets using Jira Query Language (JQL).
If no JQL is provided, will prompt for a query.`,
		Args: cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			// TODO: Implement search command
			a.logger.Info("Search command - coming soon!")
		},
	}

	cmd.Flags().IntP("limit", "l", 50, "Maximum number of results")
	cmd.Flags().Bool("json", false, "Output in JSON format")

	return cmd
}

// buildOpenCommand creates the open command.
func (a *App) buildOpenCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "open [ticket-id]",
		Short: "Open a Jira ticket in your browser",
		Long: `Open a specific Jira ticket in your default browser.
Requires a valid ticket ID (e.g., PROJ-123).`,
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			// TODO: Implement open command
			a.logger.Info("Open command - coming soon!")
		},
	}

	return cmd
}

// buildConfigCommand creates the config command.
func (a *App) buildConfigCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Manage configuration",
		Long: `Manage Jirar configuration settings.
Supports interactive setup, validation, and testing.`,
	}

	// Add subcommands
	cmd.AddCommand(a.buildConfigInitCommand())
	cmd.AddCommand(a.buildConfigShowCommand())
	cmd.AddCommand(a.buildConfigTestCommand())
	cmd.AddCommand(a.buildConfigSetCommand())

	return cmd
}

// buildConfigInitCommand creates the config init subcommand.
func (a *App) buildConfigInitCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Interactive setup wizard",
		Run: func(cmd *cobra.Command, args []string) {
			// TODO: Implement config init
			a.logger.Info("Config init command - coming soon!")
		},
	}
	return cmd
}

// buildConfigShowCommand creates the config show subcommand.
func (a *App) buildConfigShowCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show",
		Short: "Show current configuration",
		Run: func(cmd *cobra.Command, args []string) {
			// TODO: Implement config show
			a.logger.Info("Config show command - coming soon!")
		},
	}
	return cmd
}

// buildConfigTestCommand creates the config test subcommand.
func (a *App) buildConfigTestCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "test",
		Short: "Test Jira connection",
		Run: func(cmd *cobra.Command, args []string) {
			// TODO: Implement config test
			a.logger.Info("Config test command - coming soon!")
		},
	}
	return cmd
}

// buildConfigSetCommand creates the config set subcommand.
func (a *App) buildConfigSetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set [key] [value]",
		Short: "Set a specific configuration value",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			// TODO: Implement config set
			a.logger.Info("Config set command - coming soon!")
		},
	}
	return cmd
}
