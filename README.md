Установка серверной части

1. Перейди в директорию с сервером:
2. Установи зависимости Go:
- go mod tidy
3. Запусти сервер
- cd cmd/server
- go run main.go / go run .
(Сервер будет доступен по адресу ws://localhost:8080/ws)

Установка фронтенда
1. Перейди в директорию с фронтендом:
- npm install
2. Запусти фронтенд:
- npm run dev (Фронтенд будет доступен по адресу http://localhost:3000)
