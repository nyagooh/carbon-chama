FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY . .
RUN mkdir -p /app/bin && go build -o /app/bin/carbon-registry

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/bin/carbon-registry /app/
COPY frontend/ /app/frontend/

# Create directory for SQLite database
RUN mkdir -p /app/data

EXPOSE 8080
ENV PORT=8080

CMD ["./carbon-registry"]
