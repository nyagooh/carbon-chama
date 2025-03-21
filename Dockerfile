FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY . .
RUN go build -o carbon-registry

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/carbon-registry .
COPY frontend/static /app/frontend/static
COPY frontend/templates /app/frontend/templates

EXPOSE 8080
CMD ["./carbon-registry"]
