FROM golang:1.21-alpine

WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker cache
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application
COPY . .

# Build the application
RUN go build -o app

# Expose the port the app runs on
EXPOSE 1939

# Command to run the application
CMD ["./app"]
