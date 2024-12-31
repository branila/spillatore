FROM golang:1.23.3 AS builder

WORKDIR /app

COPY . .

RUN go mod tidy && go build -o SpillatoreBot .

FROM debian:bullseye-slim

WORKDIR /app

COPY --from=builder /app/SpillatoreBot /app/SpillatoreBot

COPY spillatore.json /app/spillatore.json

ENV SPILLATORE_DB=/app/spillatore.json

EXPOSE 8080

CMD ["/app/SpillatoreBot"]
