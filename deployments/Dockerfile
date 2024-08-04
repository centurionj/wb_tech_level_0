FROM golang:1.22-alpine AS builder

RUN apk --no-cache add gcc g++ libc-dev

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o main ./cmd/app

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/main /app/main
COPY templates /app/templates
COPY .env /app/.env

CMD ["./main"]
