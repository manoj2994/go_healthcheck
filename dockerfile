FROM golang:latest AS builder

WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/db_healthcheck

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/db_healthcheck .
EXPOSE 8080
CMD ["./db_healthcheck"]
