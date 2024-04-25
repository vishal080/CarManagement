# Start from the official golang image
FROM golang:1.17-alpine as builder

# Set the current working directory inside the container
WORKDIR /app

# Copy the entire source code directory into the container
COPY . .

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Build the Go app
RUN go build -o car-management.

# Start a new stage from scratch
FROM alpine:latest

# Set the working directory in the container
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/car-management.

# Expose port 8080 to the outside world
EXPOSE 8080

# Set environment variables
ENV DB_HOST=db.example.com
ENV DB_PORT=5432
ENV DB_USER=myuser
ENV DB_PASSWORD=mypassword
ENV DB_NAME=mydatabase

# Create a non-root user
RUN adduser -D -g '' appuser

# Change ownership of the application binary
RUN chown appuser:appuser /root/car-management

# Switch to the non-root user
USER appuser

# Command to run the executable
CMD ["./car-management-system"]
