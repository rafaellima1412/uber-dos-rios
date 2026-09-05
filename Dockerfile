# ---- Estágio de Build ----
FROM golang:1.24-alpine AS builder
WORKDIR /app

RUN apk add --no-cache git ca-certificates tzdata
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o /app/server ./cmd/api/main.go

# ---- Estágio de Produção ----
FROM alpine:latest
WORKDIR /app

ENV CONFIG_PATH=configs/config.yaml

# Copia binário, configs e certificados
COPY --from=builder /app/server /app/server
COPY --from=builder /app/configs /app/configs
COPY --from=builder /etc/ssl/certs /etc/ssl/certs

EXPOSE 8002

HEALTHCHECK --interval=30s --timeout=5s --retries=3 \
  CMD curl -fsS http://127.0.0.1:8002/health || exit 1

CMD ["/app/server"]
