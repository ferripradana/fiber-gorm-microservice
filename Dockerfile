FROM golang:1.20-alpine AS builder

WORKDIR /srv/go-app
COPY . .
RUN go build -o main

FROM debian:buster
WORKDIR /srv/go-app
COPY --from=builder /srv/go-app/config.json .
COPY --from=builder /srv/go-app/main .

CMD ["./main"]
