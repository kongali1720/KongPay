# Multi-stage build
FROM golang:1.21-alpine AS builder

# Install dependencies
RUN apk add --no-cache git ca-certificates tzdata

WORKDIR /build

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build with optimizations
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-w -s -X main.version=$(git describe --tags --always --dirty)" \
    -o kongpay main.go

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata curl

# Set timezone
ENV TZ=Asia/Jakarta

WORKDIR /app

# Copy binary from builder
COPY --from=builder /build/kongpay .
COPY --from=builder /build/.env .env.example

# Create non-root user
RUN adduser -D -u 1000 kongpay && chown -R kongpay:kongpay /app
USER kongpay

EXPOSE 8080

HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD curl -f http://localhost:8080/health || exit 1

CMD ["./kongpay"]
