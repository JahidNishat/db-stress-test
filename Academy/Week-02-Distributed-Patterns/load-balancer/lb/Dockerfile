FROM golang:alpine AS builder

WORKDIR /app

# Copy all go files (main.go, ring.go, health.go)
COPY *.go ./

# Initialize module and build
RUN go mod init lb
RUN go build -o lb .

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/lb .

ENTRYPOINT ["./lb"]
