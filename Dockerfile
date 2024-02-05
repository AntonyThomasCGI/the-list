# syntax=docker/dockerfile:1

FROM golang:1.18

WORKDIR /app
COPY go.mod go.sum ./
COPY . .
RUN go mod download && go mod verify
#RUN go build -v -o /usr/local/bin/the-list main.go

# Install node...
RUN apt-get update && apt-get install -y \
    software-properties-common \
    npm
RUN npm install npm@latest -g && \
    npm install n -g && \
    n latest

RUN npm ci --omit-dev
RUN npm run build

EXPOSE 8000

CMD npm run start

