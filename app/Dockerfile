# Multi-stage Dockerfile for KubeMicroServe Go application
# Stage 1: Build the Go binary
FROM golang:1.24 AS builder

# Set working directory inside the container
WORKDIR /app

# Copy Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application binary
RUN go build -o main .

# Stage 2: Create a minimal runtime image
FROM debian:bookworm-slim

# Set working directory for the runtime container
WORKDIR /app

# Copy the compiled Go binary from the builder stage
COPY --from=builder /app/main .

# Expose the application port
EXPOSE 8080

# Set the default command to run the application
CMD ["./main"]
