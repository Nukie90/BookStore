# Use the official Golang image from Docker Hub with the correct version tag
FROM golang:1.22.1

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files first (if present) and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of your application code
COPY . .

# Build your application
RUN go build -o /app/main .

# Command to run the executable
CMD ["/app/main"]
