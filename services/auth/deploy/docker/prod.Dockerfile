FROM golang:1.21-alpine3.19 as builder
WORKDIR /app

COPY ./services/auth ./
COPY ./core /core
COPY ./services/notification /services/notification
COPY ./services/planner /services/planner

RUN go mod download
RUN go build -o ./bin/auth-service ./cmd/auth-service/main.go

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/bin .
CMD ["./auth-service"]