FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o spillatore .

FROM alpine:3.14

WORKDIR /app

COPY --from=builder /app/spillatore .

EXPOSE 8080

CMD ["./spillatore"]
