RazmerRnd

Описание:

RazmerRnd — это Telegram-бот-магазин кроссовок с микросервисной архитектурой. Пользователи могут просматривать товары, добавлять в корзину и оформлять заказ. Админ может управлять каталогом.

Архитектура:

Bot Service — Telegram-бот
Catalog Service — работа с товарами
Order Service — корзина и оформление заказов
Notification Service — обработка заказов, уведомления
PostgreSQL — БД для заказов и товаров
RabbitMQ — очередь сообщений

API:

REST: GET /products, POST /orders
gRPC: ProductService, OrderService
Поддержка JSON и Protobuf

Авторизация: JWT для админов (добавление/удаление товаров)
