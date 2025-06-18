# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Копируем файлы go.mod и go.sum
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем приложение
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

# Final stage
FROM alpine:latest

WORKDIR /app

# Копируем бинарный файл из builder
COPY --from=builder /app/main .

# Запускаем приложение
CMD ["./main"]
