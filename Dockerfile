FROM golang:1.24 AS builder

WORKDIR /app
COPY . .

RUN go build -o bin/currency ./crud/cmd/currency && \
    go build -o bin/cron ./crud/cmd/cron && \
    go build -o bin/gateway ./gateway/cmd/gateway


FROM alpine

WORKDIR /app

COPY --from=builder /app/bin/* ./

CMD ["./gateway"]


