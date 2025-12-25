# Multi-stage Docker build for Jirar
FROM golang:1.21-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git ca-certificates tzdata

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags='-w -s -extldflags "-static"' \
    -a -installsuffix cgo \
    -o jirar ./cmd/jirar

# Final stage
FROM alpine:latest

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates tzdata

# Create non-root user
RUN addgroup -g 1001 -S jirar && \
    adduser -u 1001 -S jirar -G jirar

# Set working directory
WORKDIR /home/jirar

# Copy binary from builder stage
COPY --from=builder /app/jirar .

# Change ownership
RUN chown jirar:jirar /home/jirar/jirar

# Switch to non-root user
USER jirar

# Add binary to PATH
ENV PATH="/home/jirar:${PATH}"

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD jirar --version || exit 1

# Expose volumes for config
VOLUME ["/home/jirar/.jirar"]

# Set default command
ENTRYPOINT ["jirar"]
CMD ["--help"]