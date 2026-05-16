Subscription Service

REST API сервис для управления подписками пользователей.


 Стек
- Go (Gin)
- PostgreSQL
- GORM
- Docker / Docker Compose
- Swagger



Запуск проекта

```bash
docker compose up --build

Swagger

http://localhost:8080/swagger/index.html


API
Подписки

 POST /subscriptions
 GET /subscriptions
 GET /subscriptions/:id
 PUT /subscriptions/:id
 DELETE /subscriptions/:id

Агрегация

GET /subscriptions/total?user_id=UUID
(optional) service_name
(optional) start
 (optional) end

{
  "service_name": "Yandex Plus",
  "price": 400,
  "user_id": "550e8400-e29b-41d4-a716-446655440000",
  "start_date": "2025-07",
  "end_date": "2025-12"
}