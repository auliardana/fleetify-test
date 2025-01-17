# Stage 1: Builder
FROM golang:alpine AS builder

# Install git and other dependencies
RUN apk update && apk add --no-cache git
WORKDIR /app
COPY . .
RUN go build -o main ./cmd/main.go

# Stage 2: Runtime
FROM alpine
WORKDIR /app

EXPOSE 9999

# Copy the binary from the builder stage
COPY --from=builder /app/main /app

# Copy the environment file (if applicable)
COPY --from=builder /app/configs/config.yaml /app

# Run the application
CMD ["./main"]
