# Используем официальный образ Go
FROM golang:1.24-alpine AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы go.mod и go.sum
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем приложение
RUN go build -o todo-app .

# Используем легкий образ Alpine для финального контейнера
FROM alpine:latest

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем собранное приложение из builder
COPY --from=builder /app/todo-app .

# Копируем файл .env (если используется)
COPY .env .

# Открываем порт, на котором будет работать приложение
EXPOSE 3000

# Запускаем приложение
CMD ["./todo-app"]