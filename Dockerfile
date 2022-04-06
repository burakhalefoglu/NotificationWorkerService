FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/notification-worker-service .
RUN rm -rf /app/internal
RUN rm -rf /app/pkg
RUN rm -rf /app/test
RUN rm -rf /app/main.go
RUN rm -rf /app/go.mod
EXPOSE 8000

CMD [ "/app/notification-worker-service" ]