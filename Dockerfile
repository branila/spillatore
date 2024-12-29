FROM golang:1.23.3 AS builder

WORKDIR /app

COPY . .

RUN go mod tidy && go build -o SpillatoreBot .

FROM debian:bullseye-slim

# Impostiamo la directory di lavoro nel container
WORKDIR /app

# Copiamo il binario dal container di build
COPY --from=builder /app/SpillatoreBot /app/SpillatoreBot

# Copiamo il file spillatore.json nel container
COPY spillatore.json /app/spillatore.json

# Impostiamo la variabile d'ambiente per il file JSON (opzionale, dipende dal tuo codice)
ENV SPILLATORE_DB=/app/spillatore.json

# Esporre la porta su cui il bot potrebbe comunicare (se necessario)
EXPOSE 8080

# Comando per eseguire il bot
CMD ["/app/SpillatoreBot"]
