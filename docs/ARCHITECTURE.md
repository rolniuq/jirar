# Project Architecture & Structure

## Current State Analysis

### Existing Components
- **Basic Jira API Client**: Working authentication and ticket fetching
- **Configuration**: Environment-based configuration system
- **Table Display**: Basic terminal table output
- **Core Logic**: JQL query for user's tickets from last 24h

### Current Issues
- Monolithic `main.go` file (141 lines)
- No proper CLI structure
- Limited command support
- Hard-coded JQL queries
- No real-time notifications
- Basic error handling

## Target Architecture

### High-Level Architecture
```
┌─────────────────────────────────────────────────────────┐
│                    CLI Interface                        │
│  ┌─────────────┐ ┌─────────────┐ ┌─────────────┐       │
│  │    List     │ │   Search    │ │    Open     │       │
│  │   Command   │ │  Command    │ │  Command    │       │
│  └─────────────┘ └─────────────┘ └─────────────┘       │
└─────────────────────────────────────────────────────────┘
                           │
┌─────────────────────────────────────────────────────────┐
│                  Application Layer                       │
│  ┌─────────────┐ ┌─────────────┐ ┌─────────────┐       │
│  │    Config   │ │   Browser   │ │ Notification│       │
│  │  Manager    │ │  Opener     │ │  Manager    │       │
│  └─────────────┘ └─────────────┘ └─────────────┘       │
└─────────────────────────────────────────────────────────┘
                           │
┌─────────────────────────────────────────────────────────┐
│                    Service Layer                         │
│  ┌─────────────┐ ┌─────────────┐ ┌─────────────┐       │
│  │  Jira Client│ │  Formatter  │ │   Watcher   │       │
│  │             │ │             │ │             │       │
│  └─────────────┘ └─────────────┘ └─────────────┘       │
└─────────────────────────────────────────────────────────┘
                           │
┌─────────────────────────────────────────────────────────┐
│                   External APIs                          │
│  ┌─────────────┐ ┌─────────────┐ ┌─────────────┐       │
│  │ Jira REST   │ │ Desktop     │ │  System     │       │
│  │    API      │ │ Notification │ │   Browser   │       │
│  └─────────────┘ └─────────────┘ └─────────────┘       │
└─────────────────────────────────────────────────────────┘
```

## Package Structure

### `cmd/jirar/` - Entry Point
```go
// main.go
func main() {
    rootCmd := cmd.NewRootCommand()
    if err := rootCmd.Execute(); err != nil {
        os.Exit(1)
    }
}
```

### `internal/cli/` - CLI Commands
```go
// commands.go
func NewRootCommand() *cobra.Command {
    cmd := &cobra.Command{
        Use:   "jirar",
        Short: "Jira CLI tool",
        PersistentFlags: [...] // global flags
    }
    
    cmd.AddCommand(
        list.NewCommand(),
        search.NewCommand(), 
        open.NewCommand(),
        watch.NewCommand(),
        config.NewCommand(),
    )
    
    return cmd
}
```

### `internal/jira/` - Jira API Client
```go
// client.go
type Client interface {
    SearchIssues(ctx context.Context, jql string, opts ...SearchOption) (*IssueSearchResult, error)
    GetIssue(ctx context.Context, key string) (*Issue, error)
    GetCurrentUser(ctx context.Context) (*User, error)
    WatchChanges(ctx context.Context, jql string) (<-chan *Issue, error)
}

type restClient struct {
    client   *resty.Client
    baseURL  string
    email    string  
    token    string
}
```

### `internal/config/` - Configuration Management
```go
// config.go
type Config struct {
    Jira      JiraConfig      `mapstructure:"jira"`
    Notification NotificationConfig `mapstructure:"notification"`
    UI        UIConfig        `mapstructure:"ui"`
}

type JiraConfig struct {
    Domain string `mapstructure:"domain"`
    Email  string `mapstructure:"email"`
    Token  string `mapstructure:"token"`
}
```

### `internal/notification/` - Notification System
```go
// manager.go
type Manager interface {
    SendDesktopNotification(title, message string) error
    PlaySound(soundType SoundType) error
    StartWatching(ctx context.Context, filter string) (<-chan Notification, error)
}

type Notification struct {
    Issue   *jira.Issue
    Type    NotificationType
    Message string
}
```

### `internal/ui/` - Terminal UI
```go
// table.go
type TableFormatter interface {
    FormatIssues(issues []*jira.Issue, opts ...FormatOption) (string, error)
    FormatIssue(issue *jira.Issue) (string, error)
}

type JSONFormatter interface {
    FormatIssues(issues []*jira.Issue) ([]byte, error)
}
```

## Data Flow

### `jirar list` Command Flow
```
1. Parse CLI flags and arguments
2. Load configuration (file + env + flags)
3. Create Jira client with auth
4. Build JQL query from filters
5. Call Jira API
6. Process response
7. Format output (table/JSON)
8. Display to user
```

### `jirar watch` Command Flow  
```
1. Parse watch flags (interval, filters)
2. Start notification manager
3. Create Jira watcher (WebSocket/polling)
4. Listen for changes
5. Filter notifications
6. Send desktop/sound notifications
7. Handle shutdown gracefully
```

## Key Design Decisions

### 1. Interface-Based Design
- Use interfaces for all major components
- Enable easy testing and mocking
- Support multiple implementations (different notification systems)

### 2. Configuration Hierarchy
1. CLI flags (highest priority)
2. Environment variables
3. Config file
4. Default values (lowest priority)

### 3. Error Handling Strategy
- Use structured error types
- Provide user-friendly messages
- Log detailed technical errors
- Implement graceful degradation

### 4. Context Propagation
- Pass context through all layers
- Support cancellation and timeouts
- Enable distributed tracing (future)

### 5. Resource Management
- Use dependency injection
- Implement connection pooling
- Clean up resources on shutdown

## Component Interactions

### Jira Client ↔ Configuration
```go
client := jira.NewClient(config.Jira.Domain, config.Jira.Email, config.Jira.Token)
```

### Commands ↔ Jira Client
```go
issues, err := cmd.jiraClient.SearchIssues(ctx, jql, jira.WithLimit(20))
```

### UI Formatter ↔ Commands
```go
output, err := cmd.formatter.FormatIssues(issues, ui.WithColors())
fmt.Print(output)
```

### Notification Manager ↔ Watcher
```go
for change := range watcher.Changes() {
    if notificationManager.ShouldNotify(change) {
        notificationManager.SendDesktopNotification(...)
    }
}
```

## Testing Strategy

### Unit Tests
- Mock external dependencies
- Test each component in isolation
- Focus on business logic

### Integration Tests  
- Test component interactions
- Use real Jira sandbox
- Test configuration loading

### End-to-End Tests
- Test complete user workflows
- Use Docker for consistent environment
- Test error scenarios

## Migration Plan

### Phase 1: Refactor Existing Code
1. Create package structure
2. Move existing logic to appropriate packages
3. Add basic CLI structure with Cobra
4. Keep current functionality working

### Phase 2: Implement Commands
1. Implement `jirar list` command
2. Add `jirar search` command  
3. Add `jirar open` command
4. Add `jirar config` command

### Phase 3: Add Advanced Features
1. Implement notification system
2. Add `jirar watch` command
3. Enhance UI formatting
4. Add comprehensive testing

### Phase 4: Polish & Optimize
1. Performance optimization
2. Error handling improvements
3. Documentation completion
4. Build & distribution setup