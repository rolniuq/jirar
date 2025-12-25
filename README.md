# Jira CLI Tool

A modern, clean CLI tool for managing Jira tickets and receiving notifications directly from your terminal.

## Features

- 🎯 List and search Jira tickets
- 🔔 Real-time notifications
- 🚀 Fast terminal interface  
- 🎨 Beautiful table formatting
- 🔧 Flexible configuration

## Quick Start

```bash
# Build and run
go run ./cmd/jirar

# Or build binary
go build -o jirar ./cmd/jirar
./jirar --help
```

## Configuration

Set environment variables:

```bash
export JIRA_DOMAIN="your-domain.atlassian.net"
export JIRA_EMAIL="your-email@company.com"  
export JIRA_TOKEN="your-api-token"
```

Or use config file:
```bash
jirar config init  # Interactive setup
```

## Commands

```bash
jirar list          # List your tickets
jirar search "JQL"  # Search with JQL
jirar open TICKET   # Open ticket in browser
jirar watch         # Watch for notifications
jirar config        # Manage configuration
```

## Development

```bash
# Run development server
go run ./cmd/jirar

# Run tests
go test ./...

# Build for production
go build -o jirar ./cmd/jirar
```

## License

MIT License