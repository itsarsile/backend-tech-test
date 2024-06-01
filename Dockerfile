# Use official Golang image as base
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download Go module dependencies
RUN go mod download

# Copy the entire project directory to the working directory
COPY . .

# Build the Go application
RUN go build -o main .

# Expose port 8000 to the outside world
EXPOSE 8000

# Command to run the executable
CMD ["./main"]

