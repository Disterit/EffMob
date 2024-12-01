# EffMob

EffMob — это REST api для тестового задания

## Содержание
- [Особенности](#особенности)
- [Технологии](#технологии)
- [Начало работы](#начало-работы)
- [Установка](#установка)
- [API Эндпоинты](#api-эндпоинты)

## Особенности
- Добавление новых песен.
- Обновление информации о существующих песнях.
- Удаление песен из библиотеки.
- Извлечение текста песен с пагинацией по куплетам.
- Получение подробной информации о конкретных песнях.
- Извлечение полной библиотеки песен.

## Технологии
- Go (Golang)
- PostgreSQL
- gin (HTTP роутер)
- Slog (логгер)
- Swagger для документации API

## Начало работы

Чтобы запустить этот проект на своем локальном компьютере, выполните следующие шаги:

### Предварительные требования
- Go 1.18+
- Docker (опционально, для запуска PostgreSQL в контейнере)
- PostgreSQL

### Установка

1. Клонировать репозиторий:
   ```bash
   git clone https://github.com/Disterit/EffMob.git
   cd EffMob

2. Запустить приложение:
   ```bash
   docker-compose up --build

## API Эндпоинты

### song

1. **CreateSong**
   - **Endpoint:** `POST /song`
   - **Тело запроса:** JSON-объект, содержащий:
     - `song`: Название песни
     - `group`: Имя артиста/музыкальной группы
   - **Ответ:** 
     - `200 OK` при успешном добавлении песни
     - `400 Bad Request`, ошибка запроса
     - `500 Status Internal Server`, ошибка сервера

2. **GetAllSong**
   - **Эндпоинт:** `GET /song`
   - **Ответ:** 
     - `200 OK` []models.Song array of songs
     - `400 Bad Request`, ошибка запроса
     - `500 Status Internal Server`, ошибка сервера

3. **GetSongById**
   - **Эндпоинт:** `GET /song/:id`
   - **Ответ:** 
     - `200 OK` models.Song song details
     - `400 Bad Request`, ошибка запроса
     - `500 Status Internal Server`, ошибка сервера

4. **UpdateSong**
   - **Эндпоинт:** `PATCH /song/:id`
   - **Тело запроса:** JSON-объект, содержащий:
     - `name`: Название песни
     - `text`: Текст песни
     - `link`: Сслыка на песню
     - `release_date`: Дата выхода песни
   - **Ответ:** 
     - `200 OK` StatusResponse {Status: "ok"}
     - `400 Bad Request`, ошибка запроса
     - `500 Status Internal Server`, ошибка сервера
   
5. **DeleteSong**
   - **Эндпоинт:** `DELETE /song/:id`
   - **Ответ:** 
     - `200 OK` StatusResponse {Status: "ok"}
     - `400 Bad Request`, ошибка запроса
     - `500 Status Internal Server`, ошибка сервера
    
### group

1. **CreateGroup**
   - **Endpoint:** `POST /group`
   - **Тело запроса:** JSON-объект, содержащий:
     - `name`: Имя артиста/музыкальной группы
   - **Ответ:** 
     - `200 OK` map[string]interface{} id added group
     - `400 Bad Request`, ошибка запроса
     - `500 Status Internal Server`, ошибка сервера

2. **GetAllLibrary**
   - **Эндпоинт:** `GET /group`
   - **Ответ:** 
     - `200 OK` map[string][]models.Song map of groups
     - `400 Bad Request`, ошибка запроса
     - `500 Status Internal Server`, ошибка сервера

3. **GetAllSongGroupById**
   - **Эндпоинт:** `GET /group/:id`
   - **Ответ:** 
     - `200 OK` map[string][]models.Song group and they song
     - `400 Bad Request`, ошибка запроса
     - `500 Status Internal Server`, ошибка сервера

4. **UpdateGroup**
   - **Эндпоинт:** `PATCH /group/:id`
   - **Тело запроса:** JSON-объект, содержащий:
     - `name`: Имя артиста/музыкальной группы
   - **Ответ:** 
     - `200 OK` StatusResponse {Status: "ok"}
     - `400 Bad Request`, ошибка запроса
     - `500 Status Internal Server`, ошибка сервера
   
5. **DeleteGroup**
   - **Эндпоинт:** `DELETE /group/:id`
   - **Ответ:** 
     - `200 OK` StatusResponse {Status: "ok"}
     - `400 Bad Request`, ошибка запроса
     - `500 Status Internal Server`, ошибка сервера
    
### verse

1. **GetVerses**
   - **Endpoint:** `GET /song/:id/verse/:verse&limit=?`
   - **Тело запроса:** JSON-объект, содержащий:
   - **Ответ:** 
     - `200 OK` map[string]string list of verses
     - `400 Bad Request`, ошибка запроса
     - `500 Status Internal Server`, ошибка сервера

