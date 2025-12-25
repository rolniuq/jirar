# Implementation Guidelines

## Project Structure
```
jirar/
├── cmd/
│   └── jirar/                 # Main CLI entry point
│       └── main.go
├── internal/                   # Private application code
│   ├── cli/                   # CLI command handlers
│   │   ├── commands.go
│   │   ├── list.go
│   │   ├── search.go
│   │   ├── open.go
│   │   ├── watch.go
│   │   └── config.go
│   ├── jira/                  # Jira API client
│   │   ├── client.go
│   │   ├── types.go
│   │   └── auth.go
│   ├── config/                # Configuration management
│   │   ├── config.go
│   │   └── loader.go
│   ├── notification/          # Notification system
│   │   ├── desktop.go
│   │   ├── sound.go
│   │   └── manager.go
│   ├── ui/                    # Terminal UI
│   │   ├── table.go
│   │   ├── color.go
│   │   └── format.go
│   └── utils/                 # Utilities
│       ├── browser.go
│       ├── time.go
│       └── validation.go
├── pkg/                       # Public library code
│   └── jirar/                # Exported APIs if needed
├── configs/                   # Legacy configs (to be refactored)
├── docs/                      # Documentation
├── scripts/                   # Build and deployment scripts
├── go.mod
├── go.sum
├── README.md
└── Makefile
```

## Code Standards

### Go Conventions
- Follow standard Go project layout
- Use `gofmt` and `golint` for code formatting
- Package names should be short, lowercase, simple
- Interfaces should be named by the method + er suffix (e.g., `Clienter`)
- Error messages should be lowercase and not end with punctuation

### CLI Framework
- Use `cobra` for command-line interface
- Use `viper` for configuration management
- Support both flags and environment variables
- Implement proper command validation and help text

### API Design
- HTTP client should have proper timeouts and retry logic
- Support context cancellation
- Implement rate limiting for Jira API
- Use structured logging with `logrus` or `zap`

### Error Handling
- Use proper error wrapping with `fmt.Errorf` and `%w`
- Implement custom error types for different scenarios
- Return meaningful error messages to users
- Log detailed errors for debugging

## Implementation Steps

### Phase 1: CLI Foundation
1. **Setup Cobra CLI**
   - Create root command
   - Add persistent flags (verbose, config)
   - Implement version command

2. **Refactor Configuration**
   - Move from environment-only to config file support
   - Add Viper for configuration management
   - Support multiple config sources

3. **Implement Core Commands**
   - `jirar list` - Basic listing functionality
   - `jirar search` - JQL query support
   - `jirar open` - Browser opening

### Phase 2: Enhanced Features
1. **Advanced Filtering**
   - Status filtering
   - Project filtering
   - Date range filtering
   - Sorting options

2. **Output Formats**
   - Table formatting with colors
   - JSON output
   - CSV export (optional)

3. **Configuration Management**
   - Interactive setup wizard
   - Config validation
   - Connection testing

### Phase 3: Notifications
1. **Desktop Notifications**
   - Cross-platform support (macOS, Windows, Linux)
   - Custom notification content
   - Action buttons in notifications

2. **Real-time Monitoring**
   - WebSocket polling for real-time updates
   - Configurable intervals
   - Smart notification filtering

3. **Sound Alerts**
   - Different sounds for different priorities
   - Volume control
   - Mute options

### Phase 4: Advanced Features
1. **Multiple Jira Instances**
   - Profile support
   - Switching between instances
   - Instance-specific settings

2. **Ticket Management**
   - Status transitions
   - Comments
   - Assignee changes

3. **Reporting**
   - Sprint summaries
   - Time tracking
   - Productivity metrics

## Dependencies

### Required
```go
// CLI Framework
github.com/spf13/cobra v1.7.0
github.com/spf13/viper v1.15.0

// HTTP Client
github.com/go-resty/resty/v2 v2.7.0

// Terminal UI
github.com/olekukonko/tablewriter v1.1.2
github.com/fatih/color v1.15.0

// Configuration
github.com/joho/godotenv v1.5.1

// Desktop Notifications
github.com/gen2brain/beeep v0.2.4

// Browser
github.com/pkg/browser v0.0.0-20210911075727-3e5ed4a7a0e6

// Logging
github.com/sirupsen/logrus v1.9.0
```

### Optional
```go
// Testing
github.com/stretchr/testify v1.8.0
github.com/golang/mock v1.6.0

// Build
github.com/goreleaser/goreleaser v1.15.0
```

## Testing Strategy

### Unit Tests
- Test each command handler
- Test Jira API client methods
- Test configuration loading
- Test utility functions

### Integration Tests
- Test against real Jira instance (using test account)
- Test configuration file loading
- Test notification system

### End-to-End Tests
- Test complete user workflows
- Test error scenarios
- Test edge cases

## Build & Distribution

### Local Development
```bash
go run ./cmd/jirar
go build -o jirar ./cmd/jirar
./jirar --help
```

### Production Build
```bash
# Using Goreleaser
goreleaser build --snapshot
goreleaser release --clean

# Cross-compilation
GOOS=linux GOARCH=amd64 go build -o jirar-linux ./cmd/jirar
GOOS=darwin GOARCH=amd64 go build -o jirar-macos ./cmd/jirar
GOOS=windows GOARCH=amd64 go build -o jirar-windows.exe ./cmd/jirar
```

### Package Distribution
- Homebrew tap for macOS
- Chocolatey package for Windows  
- APT repository for Debian/Ubuntu
- RPM repository for Fedora/CentOS

## Performance Considerations

### API Rate Limiting
- Respect Jira API rate limits
- Implement exponential backoff
- Cache responses where appropriate

### Memory Usage
- Stream large responses
- Limit concurrent requests
- Clean up resources properly

### Startup Time
- Lazy load configuration
- Minimize import dependencies
- Cache compiled templates

## Security Considerations

### Credential Management
- Never log API tokens
- Support token encryption
- Use secure storage for credentials
- Implement token rotation

### Input Validation
- Validate JQL queries
- Sanitize user input
- Prevent command injection

### Network Security
- Use HTTPS for all API calls
- Verify SSL certificates
- Support proxy configurations