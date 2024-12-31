FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o spillatore .

FROM debian:bullseye-slim

WORKDIR /app

COPY --from=builder /app/spillatore .

EXPOSE 8080

CMD ["./spillatore"]
