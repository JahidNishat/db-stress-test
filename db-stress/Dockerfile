# Stage 1: Builder
FROM golang:alpine AS builder

WORKDIR /app

# Copy dependency files first (for caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the binary
# CGO_ENABLED=0 ensures a static binary (no C libraries needed)
RUN CGO_ENABLED=0 GOOS=linux go build -o dbstress ./cmd/dbstress

# Stage 2: Runner
FROM alpine:latest

WORKDIR /root/

# Copy binary from builder
COPY --from=builder /app/dbstress .

# Run command
ENTRYPOINT ["./dbstress"]
