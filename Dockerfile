# Use the Golang image directly
FROM golang:1.23-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go Modules and dependencies files, and download them
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app
RUN go build -o users-service main.go

# Expose port 8080
EXPOSE 8080

# Run the app
CMD ["./users-service"]
