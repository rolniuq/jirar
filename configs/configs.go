package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type JiraConfigs struct {
	Domain string
	Email  string
	Token  string
}

type AppConfig struct {
	jiraConfigs *JiraConfigs
}

func NewAppConfig() *AppConfig {
	return &AppConfig{}
}

func (a *AppConfig) Load() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	if a.jiraConfigs == nil {
		a.jiraConfigs = &JiraConfigs{
			Domain: os.Getenv("JIRA_DOMAIN"),
			Email:  os.Getenv("JIRA_EMAIL"),
			Token:  os.Getenv("JIRA_TOKEN"),
		}
	}

	return nil
}

func (a *AppConfig) GetJiraConfigs() *JiraConfigs {
	if a == nil {
		return nil
	}

	return a.jiraConfigs
}
