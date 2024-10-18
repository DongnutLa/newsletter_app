FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o newsletter ./cmd/main.go

# Crear una etapa final, más ligera, para ejecutar el binario
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/newsletter .
COPY .env .

EXPOSE 3000

# Ejecutar el binario
CMD ["./newsletter"]
