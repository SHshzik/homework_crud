## Неделя 5: Базы данных, кеширование, очереди сообщений

### Темы:
* PostgreSQL в Go: драйверы, pgx + goqu/squirell
* Redis: кеширование, pub/sub
* Kafka, RabbitMQ, NATS: обработка сообщений

### Рекомендуемые материалы:
* [Работа с SQL (Занятие 25)]
* [Очереди сообщений (Занятие 30)]
* WebSockets в Go
* [GORM](https://gorm.io/)
* [Redis](https://github.com/redis/go-redis), [Kafka](https://github.com/twmb/franz-go), [NATS](https://github.com/nats-io/nats.go), [RMQ](https://github.com/rabbitmq/amqp091-go)

### Практика:

Реализовать CRUD REST API для ресурса пользователя (user). Можно использовать наработки со второй недели. В качестве БД - PostgreSQL, для кэширования - Redis или его форк типа keyDB. Факт вызова каждого метода должен фиксироваться опциональной отправкой события в отдельный subject NATS или очередь RabbitMQ. Использовать чистую архитектуру. Данные для подключения и другие опции получать из конфигурации.
