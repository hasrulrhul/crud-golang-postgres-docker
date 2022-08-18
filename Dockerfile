FROM golang:alpine

RUN apk update && apk add --no-cache git

WORKDIR /app/src/app

COPY go.mod go.sum ./

COPY . /app/src/app

RUN go mod tidy

# RUN go build -o binary
RUN go build -o main

# Expose port 8088 to the outside world
EXPOSE 8080 8080

ENTRYPOINT ["./main"]