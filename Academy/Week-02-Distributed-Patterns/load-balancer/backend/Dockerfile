FROM golang:alpine AS builder

WORKDIR /app

# No go.mod yet in this simple folder, so we just init one on the fly or assume it exists
# For safety, let's just copy the file
COPY main.go .

# Initialize module and build
RUN go mod init backend
RUN go build -o backend main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/backend .

ENTRYPOINT ["./backend"]
