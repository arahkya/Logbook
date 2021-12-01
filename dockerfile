FROM golang:1.16-alpine

WORKDIR /src

COPY . .

RUN go mod download

COPY main.go .

RUN go build -o /app/main

RUN rm -rf *

WORKDIR /app

COPY ssl ssl
COPY resources resources

CMD ["./main"]

