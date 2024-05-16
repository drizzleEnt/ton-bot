FROM golang:1.21.1-alpine AS builder

COPY . /github.com/drizzleent/ton-bot/source/
WORKDIR /github.com/drizzleent/ton-bot/source/

RUN go mod download
RUN go build -o ./bin/crud_server cmd/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/drizzleent/ton-bot/source/.env .
COPY --from=builder /github.com/drizzleent/ton-bot/source/bin/crud_server .

EXPOSE 8080
CMD [ "./crud_server" ]