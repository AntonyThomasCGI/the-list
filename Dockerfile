# syntax=docker/dockerfile:1
FROM golang:1.18

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/the-list main.go

EXPOSE 8000

CMD ["the-list"]

