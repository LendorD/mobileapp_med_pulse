# 1. Используем официальный образ Go
FROM golang:1.24-alpine AS builder

# 2. Устанавливаем нужные пакеты
RUN apt-get update && apt-get install -y ca-certificates git ssh

# 3. Рабочая директория
WORKDIR /app

# # 4. Копируем go.mod и go.sum отдельно для кеша зависимостей
# COPY go.mod go.sum ./
# RUN --mount=type=cache,target=/go/pkg/mod \
#     go mod download

# 5. Копируем всё остальное
COPY . .

# 6. Сборка приложения
RUN --mount=type=cache,target=/go/pkg/mod \
    go build -o app ./mobileapp/cmd/app/main.go

# 7. Разрешаем запуск
RUN chmod +x app

# 9. Порт (если нужно для отладки)
EXPOSE 8080

# 10. Старт приложения
ENTRYPOINT ["./app"]
