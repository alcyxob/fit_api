# Stage 1: Build the application
FROM golang:1.23-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o trainer-app .

# Stage 2: Create a lightweight runtime image
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/trainer-app .

# Copy the SQLite database file (if it exists)
COPY trainer.db .

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["./trainer-app"]