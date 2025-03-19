#Используем Node.js 21 на базе Alpine в качестве промежуточного контейнера для сборки фронтенд-части приложения
FROM node:21-alpine AS client-builder
#Устанавливаем рабочую директорию
WORKDIR /app/client
#Копируем файлы package.json и package-lock.json для установки необходимых библиотек
COPY client/package*.json ./
#Устанавливаем необходимые библиотеки и кэшируем их
RUN npm install --frozen-lockfile
#Копируем весь фронтэнд-код в контейнер
COPY client/ . 
#Выполняем сборку клиентского приложения
RUN npm run build

#Используем Go 1.23 на базе Alpine в качестве промежуточного контейнера для сборки бэкенда
FROM golang:1.23-alpine AS server-builder

#Устанавливаем рабочую директорию
WORKDIR /app
#Копируем файлы go.mod и go.sum для загрузки необходимых библиотек Go
COPY go.mod go.sum ./
#Загружаем модули Go
RUN go mod download
#Копируем весь исходный код бэкенда в контейнер
COPY . .
#Копируем собранные файлы фронтенда в серверный контейнер
COPY --from=client-builder /app/client/dist ./client/dist
#Компилируем серверное приложение (легковесную версию, на это указывает флаг -s)
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /main

#Финальный контейнер на основе легковесного Alpine Linux
FROM alpine:3.19
#Устанавливаем рабочую директорию
WORKDIR /app
#Копируем скомпилированный бэкенд-бинарник
COPY --from=server-builder /main .
#Копируем собранные файлы фронтенда
COPY --from=client-builder /app/client/dist ./client/dist
#Копируем файлы окружения
COPY .env* ./
#Открываем порт для сервера (в нашем случае 3000)
EXPOSE ${PORT:-3000}
#Запускаем приложение
CMD [ "./main" ]