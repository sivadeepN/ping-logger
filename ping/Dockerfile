# Use the official Golang image as the base image
FROM golang:1.16-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application files to the container
COPY go.mod go.sum /app/
COPY main.go /app/

# Change directory to the app directory
WORKDIR /app

# Download Go module dependencies
RUN go mod download

# Build the Go application
RUN go build -o app

# Expose the port on which the application will listen
EXPOSE 8080

# Set the entrypoint command for the container
CMD ["./app"]
