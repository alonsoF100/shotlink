FROM golang:1.25-alpine AS builder

WORKDIR /app

# Копируем файлы зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Билдим приложение
RUN CGO_ENABLED=0 GOOS=linux go build -o shotlink ./cmd/api

# Финальный образ
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Копируем бинарник из builder stage
COPY --from=builder /app/shotlink .

# Копируем конфиги и миграции
COPY --from=builder /app/config ./config
COPY --from=builder /app/migrations ./migrations

EXPOSE 8080

CMD ["./shotlink"]