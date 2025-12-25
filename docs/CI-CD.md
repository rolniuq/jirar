# CI/CD

## GitHub Actions

This project uses GitHub Actions for continuous integration and deployment.

### Workflows

#### **CI/CD Pipeline** (`.github/workflows/ci-cd.yml`)
- **Triggers**: Push to main/develop, PRs, releases
- **Jobs**:
  - **test**: Unit tests, linting, code coverage
  - **security**: Security scanning (gosec, govulncheck)
  - **build**: Multi-platform builds (Linux, macOS, Windows)
  - **release-build**: Automated releases via Goreleaser
  - **docker**: Multi-arch Docker builds
  - **notify**: Slack notifications on failure

#### **Dependency Updates** (`.github/workflows/dependency-update.yml`)
- **Triggers**: Daily schedule, manual dispatch
- **Features**: Automated dependency updates with PRs

#### **Quality Checks** (`.github/workflows/quality-checks.yml`)
- **Coverage analysis**: 50% minimum coverage requirement
- **Code quality**: golangci-lint, gofmt, goimports
- **Complexity analysis**: Cyclomatic complexity limits

#### **Quick Test** (`.github/workflows/quick-test.yml`)
- **Purpose**: Fast validation on every push
- **Features**: Build and CLI functionality tests

### Local Development

Use the Makefile for local development:

```bash
# Set up development environment
make setup

# Run all CI checks locally
make ci

# Build for current platform
make build

# Build for all platforms
make build-all

# Run tests with coverage
make test-coverage

# Format code
make fmt

# Run linting
make lint

# Security checks
make security
```

### Code Quality Standards

#### **Linting Rules** (`.golangci.yml`)
- Complexity limit: 10
- All major Go linters enabled
- Custom exclusions for test files
- Strict formatting requirements

#### **Testing Requirements**
- Minimum 50% code coverage
- Race condition testing enabled
- Integration tests for critical paths
- Benchmark tests for performance monitoring

#### **Security Standards**
- No hardcoded credentials
- Dependency vulnerability scanning
- SAST (Static Analysis Security Testing)
- Container security best practices

### Release Process

#### **Automated Releases**
1. Create a new tag: `git tag v1.0.0`
2. Push the tag: `git push origin v1.0.0`
3. GitHub Actions will:
   - Build binaries for all platforms
   - Create GitHub release
   - Build and push Docker images
   - Update Homebrew formula (if configured)

#### **Release Artifacts**
- **Binaries**: Linux (amd64/arm64), macOS (amd64/arm64), Windows (amd64)
- **Docker**: Multi-arch images for Linux
- **Documentation**: Auto-generated from code comments

### Environment Variables

#### **Required for CI/CD**
```yaml
# GitHub Secrets
GITHUB_TOKEN:          # GitHub API token
DOCKERHUB_USERNAME:   # Docker Hub username
DOCKERHUB_TOKEN:      # Docker Hub access token
SLACK_WEBHOOK_URL:    # Slack notifications (optional)
```

#### **Required for Local Development**
```bash
# Jira Configuration
JIRA_DOMAIN=your-domain.atlassian.net
JIRA_EMAIL=your-email@company.com
JIRA_TOKEN=your-api-token

# Optional
JIRAR_DEBUG=false
JIRAR_LOG_LEVEL=info
```

### Performance Monitoring

#### **Metrics Collected**
- Build time trends
- Test coverage over time
- Code quality scores
- Security vulnerability count

#### **Alerting**
- Build failures → Slack notifications
- Security vulnerabilities → GitHub issues
- Performance degradation → PR reviews

### Containerization

#### **Docker Support**
- Multi-stage builds for small images
- Security hardening (non-root user)
- Health checks included
- ARM64 support for Apple Silicon

#### **Docker Commands**
```bash
# Build locally
make docker-build

# Run locally
make docker-run

# Pull from registry
docker pull jirar/jirar:latest
```

### Contributing to CI/CD

#### **Adding New Checks**
1. Update `.github/workflows/ci-cd.yml`
2. Add corresponding Makefile targets
3. Update documentation

#### **Quality Standards**
- All checks must pass before merge
- New features need test coverage
- Security review for sensitive changes
- Performance impact assessment for heavy operations

### Troubleshooting

#### **Common Issues**
- **Go version mismatch**: Use Go 1.21+
- **Dependency conflicts**: Run `make deps-update`
- **Linting failures**: Run `make fmt` and `make lint`
- **Test failures**: Check environment variables

#### **Debug Mode**
```bash
# Enable debug logging
JIRAR_DEBUG=true ./dist/jirar --log-level debug

# Run with verbose output
make ci VERBOSE=true
```