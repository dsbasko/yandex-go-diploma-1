FROM golang:1.21-alpine3.19 as builder
WORKDIR /app

COPY ./services/notification ./
COPY ./core /core
COPY ./services/auth /services/auth
COPY ./services/planner /services/planner

RUN go mod download
RUN go build -o ./bin/notification-service ./cmd/notification-service/main.go

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/bin .
CMD ["./notification-service"]