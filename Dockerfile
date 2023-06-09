# Use the official Go image as the base image
FROM golang:1.16

# Set the working directory inside the container
WORKDIR /app

# Copy the Go server code into the container
COPY . .

# Build the Go server
RUN go build

# Set the command to run the server when the container starts
CMD ["./code-snippet-manager"]

# Expose the port that the server listens on
EXPOSE 8080