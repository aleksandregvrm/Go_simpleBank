# Build stage
FROM golang:1.23-alpine3.21 AS builder

WORKDIR /app

COPY . .

# Install necessary tools and build the Go application
RUN apk add --no-cache curl
RUN go build -o main main.go

# Download and install the migrate binary
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar -xz \
    && mv migrate /usr/local/bin/migrate

# Runtime stage
FROM alpine:3.21

WORKDIR /app
 
# Copy the built application and necessary files
COPY --from=builder /app/main .
COPY --from=builder /usr/local/bin/migrate /usr/local/bin/migrate
COPY .env .
COPY start.sh .
RUN chmod +x start.sh 
COPY db/migration ./migration

# Expose the application port
EXPOSE 8080

# Use the entrypoint script for migrations and starting the app
ENTRYPOINT ["/app/start.sh"]
CMD ["/app/main"]
