---
title: README
tags: []
---
# **Служба каталогов автомобилей**

Этот проект реализует RESTful API для управления каталогом автомобилей. Он включает функции для выборки, добавления, обновления и удаления записей об автомобилях с поддержкой фильтрации и разбивки на страницы. Кроме того, он интегрируется с внешним API для расширения информации об автомобиле при добавлении. Сервис поддерживается базой данных PostgreSQL с возможностью миграции для управления схемой базы данных.

## **Характеристики**

- **RESTful API**\: реализует операции CRUD для управления каталогом автомобилей.
- **Фильтрация и разбивка на страницы**\: поддерживает выборку данных об автомобиле с помощью фильтрации и разбивки на страницы.
- **Интеграция с внешним API**\: обогащает информацию об автомобиле за счет извлечения данных из внешнего API.
- **Управление базами данных**\: Использует PostgreSQL для хранения данных с миграциями для управления схемами.
- **Ведение журнала**\: включает в себя ведение журнала отладки и информации для лучшей наблюдаемости.
- **Конфигурация**\: сохраняет конфигурацию в `.env` файле для упрощения настройки.
- **Документация Swagger**\: создает документацию Swagger для API.

## **Приступая к работе**

### **Предварительные требования**

- Перейти на версию Golang 1.16 или более позднюю
- PostgreSQL 13 или более поздней версии

### **Установка**

**Клонировать репозиторий**\:

```
git clone https://github.com/zatrasz75/tz_go.git
cd tz_go
```

**Настройка переменных окружения**\:

Скопируйте `.env.example` файл в `.env` и обновите переменные по мере необходимости:

```
APP_IP="0.0.0.0"
APP_PORT="4141"

POSTGRES_USER="postgres"
POSTGRES_PASSWORD="postgrespw"
POSTGRES_DB="clean-tz"
URL_DB="tz-db"
PORT_DB="5432"

# Заменить на url внешнего api для обагащение добавленных данных
EXTERNAL_API_URL="https://api.agify.io/?name="
```

Отредактируйте `.env` файл, чтобы задать строку подключения к PostgreSQL и URL внешнего API.

**Настройка Миграции**\:

Файлы для миграции находятся в директории migrations

```
# Создать файл миграции

sql-migrate new <имя файла>

# и отредактировать 
```

**Запустите сервис**\:

```
go run ./cmd/main.go
```

Это запустит сервер с IP-адресом и портом, указанными в `.env` файле.

### **Документация по API**

Запуск сервера на <http://localhost:4242>

Документация Swagger API: <http://localhost:4242/swagger/index.html>

### **Или установка в Docker**

```
docker compose up -d
```

### **Документация по API**

Запуск сервера на <http://localhost:4242>

Документация Swagger API: <http://localhost:4242/swagger/index.html>

## **Использование**

### **Добавление Автомобилей**

Чтобы добавить новые автомобили, отправьте запрос POST на `/cars` со следующей полезной нагрузкой в формате JSON:

```
{
   "regNums": ["X123XX150"]
}
```

Сервис обогатит информацию об автомобиле, отправив запрос во внешний API, а затем сохранит обогащенные данные в базе данных.

### **Подбор автомобилей**

Чтобы извлекать автомобили с фильтрацией и разбиением на страницы, используйте `/cars` конечную точку с параметрами запроса для фильтрации и разбиения на страницы.

### **Обновление автомобилей**

Чтобы обновить информацию об автомобиле, используйте `/cars/{id}` конечную точку с запросом PUT и обновленными данными об автомобиле.

### **Удаление автомобилей**

Чтобы удалить car, используйте `/cars/{id}` конечную точку с запросом на УДАЛЕНИЕ.