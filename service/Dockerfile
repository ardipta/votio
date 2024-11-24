# Service Dockerfile
# Service Dockerfile
FROM golang:1.23-alpine

WORKDIR /app

# Install Air
RUN go install github.com/air-verse/air@latest

# Copy dependencies first for caching
COPY service/go.mod service/go.sum ./
RUN go mod download

# Copy application files and .air.toml explicitly
COPY service/ .
COPY .air.toml /app/.air.toml

# Use Air with the configuration file
CMD ["air", "-c", "/app/.air.toml"]

