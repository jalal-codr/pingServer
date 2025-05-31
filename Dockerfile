# Start from the official Go image
FROM golang:1.21-alpine

# Set environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Install git (required for Go module fetching), ca-certificates for HTTPS
RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project
COPY . .

# Build the Go binary
RUN go build -o app .

# Expose the port your app listens on
EXPOSE 7080

# Run the app
CMD ["./app"]
