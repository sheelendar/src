# Dockerfile

# Use the official Go image as the base image
FROM golang:1.20.5-alpine

# Set the working directory inside the container
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
# Copy the Go application code into the container

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o /worker

# Command to run the Go application
CMD ["/worker"]