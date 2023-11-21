# Используем образ с Go для сборки
FROM golang:1.21.3-alpine AS build

# Установим рабочую директорию /app
WORKDIR /app

# Скопируем все файлы из текущего каталога в /app внутри контейнера
COPY . .

# Загрузим зависимости Go
RUN go mod download

# Соберем исполняемый файл
RUN go build -o gRPCApiGo ./cmd/server/main.go

# Второй этап - запустим приложение
FROM alpine:latest

# Скопируем исполняемый файл из предыдущего образа
COPY --from=build /app/gRPCApiGo .
COPY --from=build /app/migrations  /migrations

# Откроем порт, на котором будет работать приложение
EXPOSE 50051

# Запустим приложение при старте контейнера
CMD ["./gRPCApiGo"]