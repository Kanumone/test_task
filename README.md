# Avito test task

## Для запуска необходимо:
1. Установить Docker compose
2. Установить Postman
3. Создать .env файл и заполнить параметры доступа к БД
```bash
make env
```
4. Запустить контейнер
```bash
make up_container
```
5. При необходимости изменить переменные окружения и запустить сервер
```bash
make run
```

## Тестирование и документация

### Методы API:
1. POST http://localhost:8080/slug

```bash
curl --location 'http://localhost:8080/slug' \
--data '{
    "title": "AVITO_VOICE_MESSAGES"
}'

```

2. DELETE http://localhost:8080/slug

```bash
curl -X --location 'http://localhost:8080/slug' \
--data '{
    "title": "AVITO_VOICE_MESSAGES"
}'

```

3. GET http://localhost:8080/user

```bash
curl -X GET --location 'http://localhost:8080/user' \
--data '{
    "user_id": 1
}'

```

4. PATCH http://localhost:8080/user

```bash
curl -X PATCH --location 'http://localhost:8080/user' \
--data '{
    "user_id": 1,
    "add_slugs": ["AVITO_VALID_SLUG", "AVITO_VOICE"],
    "delete_slugs": ["AVITO_VALID_SLUG", "AVITO_VOICE"]
}'

```
### Потестировать можно с помощью POSTMAN

[![Run in Postman](https://run.pstmn.io/button.svg)](https://www.postman.com/lunar-module-explorer-62645013/workspace/kanumone/overview)
