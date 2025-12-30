# Build stage
FROM golang:1.24-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Run stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates wget

# Add non-root user
RUN addgroup -g 1001 -S appuser && \
    adduser -u 1001 -S appuser -G appuser

WORKDIR /home/appuser

# Copy the binary from builder
COPY --from=builder --chown=appuser:appuser /app/main .

# Copy runtime dependencies
COPY --from=builder --chown=appuser:appuser /app/migrations ./migrations
COPY --from=builder --chown=appuser:appuser /app/rbac/rbac_model.conf ./rbac/rbac_model.conf

USER appuser

EXPOSE 8080

HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD wget --no-verbose --tries=1 --spider http://localhost:8080/health || exit 1

CMD ["./main"]
