// Package config handles application configuration using Viper.
package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config represents the application configuration.
type Config struct {
	Jira     JiraConfig `mapstructure:"jira"`
	UI       UIConfig   `mapstructure:"ui"`
	Debug    bool       `mapstructure:"debug"`
	LogLevel string     `mapstructure:"log_level"`
}

// JiraConfig holds Jira-specific configuration.
type JiraConfig struct {
	Domain string `mapstructure:"domain"`
	Email  string `mapstructure:"email"`
	Token  string `mapstructure:"token"`
}

// UIConfig holds UI-specific configuration.
type UIConfig struct {
	Colors  bool `mapstructure:"colors"`
	Icons   bool `mapstructure:"icons"`
	Compact bool `mapstructure:"compact"`
}

// New creates a new configuration instance with defaults.
func New() (*Config, error) {
	cfg := &Config{}

	// Set defaults
	setDefaults()

	// Configure Viper
	viper.AutomaticEnv()
	viper.SetEnvPrefix("JIRAR")

	// Environment variable mappings
	viper.BindEnv("jira.domain", "JIRA_DOMAIN")
	viper.BindEnv("jira.email", "JIRA_EMAIL")
	viper.BindEnv("jira.token", "JIRA_TOKEN")

	// Load configuration file
	if err := loadConfigFile(); err != nil {
		// Config file is optional, but log warning if not found
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			fmt.Printf("Warning: %v\n", err)
		}
	}

	// Unmarshal into struct
	if err := viper.Unmarshal(cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	// Validate configuration
	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("configuration validation failed: %w", err)
	}

	return cfg, nil
}

// setDefaults sets default values for configuration.
func setDefaults() {
	viper.SetDefault("debug", false)
	viper.SetDefault("log_level", "info")
	viper.SetDefault("ui.colors", true)
	viper.SetDefault("ui.icons", true)
	viper.SetDefault("ui.compact", false)
}

// loadConfigFile loads configuration from various locations.
func loadConfigFile() error {
	// Try to find config file in standard locations
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	// Search paths in order of preference
	configPaths := []string{
		".",
		"$HOME/.jirar",
		"$HOME/.config/jirar",
		"/etc/jirar",
	}

	for _, path := range configPaths {
		viper.AddConfigPath(path)
	}

	return viper.ReadInConfig()
}

// Validate validates the configuration.
func (c *Config) Validate() error {
	if c.Jira.Domain == "" {
		return fmt.Errorf("jira domain is required")
	}
	if c.Jira.Email == "" {
		return fmt.Errorf("jira email is required")
	}
	if c.Jira.Token == "" {
		return fmt.Errorf("jira token is required")
	}
	return nil
}

// IsDebug returns true if debug mode is enabled.
func (c *Config) IsDebug() bool {
	return c.Debug
}

// GetLogLevel returns the configured log level.
func (c *Config) GetLogLevel() string {
	return c.LogLevel
}
