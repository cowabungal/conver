FROM golang:1.16-alpine AS builder

RUN go version

COPY . /conver/
WORKDIR /conver/

RUN go mod download
RUN GOOS=linux go build -o ./.bin/bot ./cmd/bot/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /conver/.bin/bot .
COPY --from=0 /conver/ .

CMD ["./bot"]
