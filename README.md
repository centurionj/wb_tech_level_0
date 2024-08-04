## Обзор

Тестовое задание Wildberries

## Задание

В БД:

- Развернуть локально postgresql
- Создать свою бд
- Настроить своего пользователя.
- Создать таблицы для хранения полученных данных.

В сервисе:

1. Подключение и подписка на канал в nats-streaming
2. Полученные данные писать в Postgres
3. Так же полученные данные сохранить in memory в сервисе (Кеш)
4. В случае падения сервиса восстанавливать Кеш из Postgres
5. Поднять http сервер и выдавать данные по id из кеша
6. Сделать простейший интерфейс отображения полученных данных, для
   их запроса по id

Доп инфо:

- Данные статичны, исходя из этого подумайте насчет модели хранения в Кеше и в pg. Модель в файле model.json
- В канал могут закинуть что угодно, подумайте как избежать проблем из-за этого
- Чтобы проверить работает ли подписка онлайн, сделайте себе отдельный скрипт, для публикации данных в канал
- Подумайте как не терять данные в случае ошибок или проблем с сервисом
- Nats-streaming разверните локально ( не путать с Nats )

Модель (model.json):

```azure
{
"order_uid": "b563feb7b2b84b6test",
"track_number": "WBILMTESTTRACK",
"entry": "WBIL",
"delivery": {
"name": "Test Testov",
"phone": "+9720000000",
"zip": "2639809",
"city": "Kiryat Mozkin",
"address": "Ploshad Mira 15",
"region": "Kraiot",
"email": "test@gmail.com"
  },
"payment": {
"transaction": "b563feb7b2b84b6test",
"request_id": "",
"currency": "USD",
"provider": "wbpay",
"amount": 1817,
"payment_dt": 1637907727,
"bank": "alpha",
"delivery_cost": 1500,
"goods_total": 317,
"custom_fee": 0
  },
"items": [
    {
      "chrt_id": 9934930,
      "track_number": "WBILMTESTTRACK",
      "price": 453,
      "rid": "ab4219087a764ae0btest",
      "name": "Mascaras",
      "sale": 30,
      "size": "0",
      "total_price": 317,
      "nm_id": 2389212,
      "brand": "Vivienne Sabo",
      "status": 202
    }
  ],
"locale": "en",
"internal_signature": "",
"customer_id": "test",
"delivery_service": "meest",
"shardkey": "9",
"sm_id": 99,
"date_created": "2021-11-26T06:22:19Z",
"oof_shard": "1"
}
```

### Начало работы

В корневой директории создать .env из .env.example

### Параметры приложения

| Ключ                | Значение                               | Примечания                                                                 |
|---------------------|----------------------------------------|----------------------------------------------------------------------------|
| `GIN_MODE`          | Указываем какие настройки использовать | *                                                                          |
| `TIME_ZONE`         | Временная зона                         | *                                                                          |
| `POSTGRES_DB`       | Имя базы данных postgres               | Должен быть отличный от postgres                                           |
| `POSTGRES_USER`     | Имя пользователя базы данных postgres  | Должен быть отличный от postgres                                           |
| `POSTGRES_PASSWORD` | Пароль к базе данных postgres          | Должен быть отличный от postgres                                           |
| `POSTGRES_HOST`     | Хост базы данных postgres              | Если запуск через докер, то ставим "db", ручками - "localhost"             |
| `POSTGRES_PORT`     | Порт базы данных postgres              | *                                                                          |
| `SSL_MODE`          | SSL к базе данных postgres             | *                                                                          |
| `REDIS_HOST`        | Хост redis                             | Если запуск через докер, то ставим "redis", ручками - "localhost"          |
| `REDIS_PORT`        | Порт redis                             | *                                                                          |
| `NATS_HOST`         | Хост stan'a                            | Если запуск через докер, то ставим "nats-streaming", ручками - "localhost" |
| `NATS_PORT`         | Порт stan'a                            | *                                                                          |
| `HTTP_PORT`         | Порт дл запуска web-приложения         | *                                                                          |

### Запуск всего приложения через docker & docker-compose

В корне проекта
```shell
docker-compose up --build
```
Приложение доступно по адресу: http://localhost/

### Локальный запуск

1. Установка зависимостей
```shell
go mod download
```

2. Поднимаем базу кэш и стан
```shell
docker-compose up db redis nats-streaming
```

3. Запуск web-приложения
```shell
go run ./cmd/app/main.go
```
Приложение доступно по адресу: http://localhost:8080/

### Запуск скрипта для публикации данных в канал

В корне проекта

```shell
go run ./scripts/publish.go
```
- P.S. при повторном запуске скрипта в файле /scripts/publish.go необходимо менять поле "OrderUID".
