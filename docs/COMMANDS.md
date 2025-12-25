# Jira CLI - Commands Specification

## Overview
Jirar is a command-line tool for managing Jira tickets and receiving notifications directly from your terminal.

## Global Options
```
--config, -c     Path to config file (default: ~/.jirar/config.json)
--verbose, -v    Enable verbose logging
--help, -h       Show help
--version        Show version
```

## Commands

### `jirar list`
List all tickets currently assigned to you.

**Usage:**
```bash
jirar list [options]
```

**Options:**
```
--status, -s     Filter by status (todo, in-progress, done, resolved)
--limit, -l      Maximum number of tickets to show (default: 20)
--project, -p    Filter by project key
--sort           Sort field (updated, created, priority) (default: updated)
--json           Output in JSON format
```

**Examples:**
```bash
jirar list                           # List all my tickets
jirar list --status todo             # Only TODO tickets
jirar list --project PROJ --limit 10 # 10 tickets from PROJ project
jirar list --json                    # Output as JSON
```

### `jirar search`
Search tickets using custom JQL query.

**Usage:**
```bash
jirar search "<jql-query>" [options]
```

**Options:**
```
--limit, -l      Maximum results (default: 50)
--json           Output in JSON format
```

**Examples:**
```bash
jirar search "project = PROJ AND status = 'In Progress'"
jirar search "assignee = currentUser() AND updated >= -7d"
jirar search "priority = Highest" --limit 5
```

### `jirar open`
Open a specific Jira ticket in your default browser.

**Usage:**
```bash
jirar open <ticket-id>
```

**Examples:**
```bash
jirar open PROJ-123
jirar open TICKET-456
```

### `jirar watch`
Start monitoring for real-time notifications.

**Usage:**
```bash
jirar watch [options]
```

**Options:**
```
--interval, -i   Polling interval in seconds (default: 30)
--filter         Custom JQL filter for notifications
--sound          Enable sound notifications
--desktop        Enable desktop notifications
```

**Examples:**
```bash
jirar watch                           # Start watching with defaults
jirar watch --interval 60 --desktop   # Check every 60s with desktop notifications
jirar watch --filter "priority = Highest"  # Only high priority tickets
```

### `jirar config`
Setup and manage configuration.

**Usage:**
```bash
jirar config [command] [options]
```

**Subcommands:**
```
init            Interactive setup wizard
show            Show current configuration
test            Test connection to Jira
set <key> <value>  Set specific config value
```

**Examples:**
```bash
jirar config init                     # Interactive setup
jirar config show                     # Show current config
jirar config test                     # Test Jira connection
jirar config set domain company.atlassian.net
```

## Output Formats

### Table Format (Default)
```
┌─────────┬────────────┬──────────┬──────────────────────┬─────────────────────────────┐
│ Key     │ Status     │ Updated  │ Summary               │ Link                        │
├─────────┼────────────┼──────────┼──────────────────────┼─────────────────────────────┤
│ PROJ-123│ 🔥 In Progress │ 14:30 25/12 │ Fix login bug         │ https://company.atlassian.net/browse/PROJ-123 │
└─────────┴────────────┴──────────┴──────────────────────┴─────────────────────────────┘
```

### JSON Format
```json
{
  "issues": [
    {
      "key": "PROJ-123",
      "fields": {
        "summary": "Fix login bug",
        "status": {
          "name": "In Progress"
        },
        "updated": "2023-12-25T14:30:00.000+0000"
      },
      "url": "https://company.atlassian.net/browse/PROJ-123"
    }
  ],
  "total": 1,
  "start": 0,
  "limit": 20
}
```

## Status Indicators
- `🔥` In Progress
- `✅` Done/Resolved
- `📋` To Do
- `⏸️` Blocked
- `👀` Review

## Configuration Priority
1. Command line flags
2. Environment variables
3. Config file (~/.jirar/config.json)
4. Default values

## Environment Variables
```
JIRA_DOMAIN      Jira instance URL
JIRA_EMAIL       User email for authentication
JIRA_TOKEN       API token
JIRA_CONFIG_PATH Path to config file
```