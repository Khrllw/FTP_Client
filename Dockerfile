FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
COPY .env /app/.env

RUN go build -o ftp-client ./cmd/app

FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/ftp-client .
COPY --from=builder /app/docs ./docs
COPY --from=builder /app/.env /app/.env

EXPOSE 8080

CMD ["./ftp-client"]
