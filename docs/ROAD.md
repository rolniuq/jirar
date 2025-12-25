# Jira CLI Tool Roadmap

## Problem Statement
As an engineer, opening a browser just to check Jira notifications is inefficient. I want desktop notifications and quick CLI access to Jira tickets assigned to me.

## Current Status ✅
- Basic Jira API integration working
- Fetches tickets updated in last 24h where I'm assignee/watcher/mentioned
- Displays results in formatted table
- Environment-based configuration

## Roadmap

### Phase 1: Core CLI Features (Current)
- [x] Basic Jira API authentication
- [x] Fetch personal tickets from last 24h
- [x] Display results in table format
- [x] Environment configuration

### Phase 2: Enhanced CLI Commands
- [ ] `jirar list` - List all my assigned tickets
- [ ] `jirar watch` - Start watching for real-time notifications
- [ ] `jirar search <query>` - Search tickets with custom JQL
- [ ] `jirar config` - Setup/configuration wizard
- [ ] `jirar open <ticket-id>` - Open ticket in browser

### Phase 3: Notification System
- [ ] Desktop notifications (macOS/Windows/Linux)
- [ ] WebSocket/long-polling for real-time updates
- [ ] Notification filtering rules
- [ ] Sound alerts for urgent tickets
- [ ] Configurable notification channels (Slack, Teams)

### Phase 4: Advanced Features
- [ ] Multiple Jira instance support
- [ ] Ticket status transitions from CLI
- [ ] Time tracking integration
- [ ] Sprint/team views
- [ ] Dashboard/summary reports

### Phase 5: Distribution & UX
- [ ] Binary releases for major platforms
- [ ] Homebrew package
- [ ] Auto-update mechanism
- [ ] Rich terminal UI (fancy tables, colors)
- [ ] Shell completion scripts

## Technical Architecture
```
┌─────────────┐    ┌──────────────┐    ┌─────────────┐
│   CLI       │────│   Core API   │────│  Jira API   │
│ Commands    │    │ Service      │    │ Integration │
└─────────────┘    └──────────────┘    └─────────────┘
                           │
                 ┌──────────────┐
                 │ Notification │
                 │   Manager    │
                 └──────────────┘
```

## Key Design Decisions
- **CLI-first**: Focus on terminal experience, GUI secondary
- **Real-time**: Push notifications vs polling
- **Multi-platform**: macOS, Windows, Linux support
- **Privacy-first**: Local storage, no cloud services
- **Extensible**: Plugin architecture for custom integrations
