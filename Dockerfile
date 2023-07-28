# Use the official Go image as the base image
FROM golang:1.19

# Set the working directory inside the container
WORKDIR /usr/app

# Copy the entire Go backend project to the container's working directory
COPY . .

# Build the Go backend
RUN go mod download && go mod verify
RUN GOOS=linux GOARCH=amd64 go build -v -o main

# Expose the port that the Go backend is listening on
EXPOSE 8080

# Start the Go backend server
CMD ["./main"]