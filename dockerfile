# Use the official Golang image from Docker Hub with the correct version tag
FROM golang:1.22.1

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files first (if present) and download dependencies
COPY . .

RUN go get -d -v ./...

RUN go build -o main .

# Expose the port the app runs on
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
