FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker cache
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application
COPY . .

# Build the application
RUN mkdir -p /app/bin && go build main.go -o /app/bin/carbon-registry

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/bin/carbon-registry /app/
COPY frontend/ /app/frontend/

# Create directory for SQLite database
RUN mkdir -p /app/data

# Expose the port the app runs on
EXPOSE 8080
ENV PORT=8080

# Command to run the application
CMD ["./app/carbon-registry"]
