# Use the official Golang image as a base image
FROM golang:1.20-alpine as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the go.mod file (excluding go.sum since it's not present)
COPY go.mod ./

# Download all dependencies. Dependencies will be cached if the go.mod file is not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o taxi-fare

# Start a new stage from scratch
FROM alpine:latest

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/taxi-fare .

# Copy the input file
COPY input.txt .

# Command to run the executable
CMD ["./taxi-fare"]
