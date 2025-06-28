# Build stage
FROM golang:1.23-alpine AS builder

# Build arguments
ARG ENV=dev
ARG VERSION=latest

# Install git (required for some Go modules)
RUN apk add --no-cache git

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o auth-service ./cmd/app

# Final stage
FROM alpine:latest

# Build arguments (pass from build stage)
ARG ENV=dev
ARG VERSION=latest

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

# Create non-root user for security
RUN adduser -D -s /bin/sh appuser

WORKDIR /root/

# Copy the binary from builder stage
COPY --from=builder /app/auth-service .

# Copy config directory
COPY --from=builder /app/configs ./configs

# Set environment variables
ENV APP_ENV=${ENV}
ENV APP_VERSION=${VERSION}

# Change ownership to appuser
RUN chown -R appuser:appuser /root/

# Switch to non-root user
USER appuser

# Expose port (default from your config is 8080)
EXPOSE 8080

# Command to run the executable
CMD ["./auth-service"] 