FROM golang:1.21-alpine3.19
WORKDIR /app
CMD ["go", "run", "./cmd/planner-service/main.go"]