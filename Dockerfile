FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o spillatore .

FROM alpine:3.14

WORKDIR /app

# Installa glibc per compatibilità
RUN apk add --no-cache libc6-compat

COPY --from=builder /app/spillatore .

EXPOSE 8080

CMD ["./spillatore"]
