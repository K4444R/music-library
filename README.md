# Решение задания на вакансию Junior GO Developer

## Описание задачи

Необходимо реализовать RESTful API для управления онлайн-библиотекой песен. Задание включает следующие требования:

1. **Реализовать REST-методы**:
   - Получение данных библиотеки с фильтрацией по всем полям и пагинацией.
   - Получение текста песни с пагинацией по куплетам.
   - Удаление песни.
   - Изменение данных песни.
   - Добавление новой песни в формате JSON.

2. **Интеграция с внешним API**:
   - При добавлении песни выполняется запрос к внешнему API для получения дополнительной информации (дата выпуска, текст песни, ссылка).
   - Внешний API описан в Swagger и будет предоставлен при проверке задания. Реализовывать его отдельно не нужно.

3. **База данных**:
   - Использовать PostgreSQL для хранения данных.
   - Структура базы данных должна создаваться автоматически при старте сервиса с помощью миграций.

4. **Логирование**:
   - Код должен быть покрыт debug- и info-логами для отслеживания выполнения операций.

5. **Конфигурация**:
   - Все конфигурационные данные (например, подключение к базе данных, API-ключи) должны быть вынесены в `.env`-файл.

6. **Документация**:
   - Сгенерировать Swagger-документацию для реализованного API.

---

## Реализация

### 1. REST-методы

Реализованы следующие методы:

- **GET /songs** — получение списка песен с фильтрацией и пагинацией.
- **POST /songs** — добавление новой песни.
- **PUT /songs/:id** — изменение данных песни.
- **DELETE /songs/:id** — удаление песни.

Пример запроса на добавление песни:
```json
{
    "group": "Muse",
    "song": "Supermassive Black Hole"
}
```

### 2. Интеграция с внешним API

При добавлении песни выполняется запрос к внешнему API для получения дополнительной информации:
- **Дата выпуска** (`releaseDate`)
- **Текст песни** (`text`)
- **Ссылка** (`link`)

Пример ответа от внешнего API:
```json
{
"releaseDate": "08 Jul 2009, 00:18",
    "text": "My Fairy King\", written by Freddie Mercury, for Queen's eponymous first album, deals with Rhye, a fantasy world created by Mercury as a child and featured in other Queen songs, most notably \"Seven Seas of Rhye\". \n\n\"My Fairy King\" is the first song on the album to feature Mercury's piano skills – as the piano on \"Doing All Right\" was played by guitarist Brian May. The guitarist  was quite impressed by Mercury's piano playing on the track, and from this point on Mercury handled most of Queen's piano parts.\n\nBefore writing this song Mercury was known as Freddie Bulsara, and this song is said to have inspired him to change his surname. Its lyrics contain a verse with the words \"Mother Mercury, look what they've done to me.\" Brian May has said that after the line was written, Freddie claimed he was singing about his mother. Subsequently, Freddie Bulsara took the stage name Freddie Mercury. This was another attempt to separate his stage persona (\"extroverted monster\", as Mercury himself once described it) from his personal persona (introverted).\n\nWritten during the band's time in the studio, the song contains many voice overdubs and vocal harmonies, which Mercury was fond of. Drummer Roger Taylor also displays his vocal skills here, hitting some of the highest notes in the composition. The vocal overdubs technique would later be used in many Queen songs, most notably \"Bohemian Rhapsody\". \nMercury borrowed some lines from Robert Browning's poem, \"The Pied Piper of Hamelin\". <a href=\"http://www.last.fm/music/Queen/_/My+Fairy+King\">Read more on Last.fm</a>. User-contributed text is available under the Creative Commons By-SA License; additional terms may apply.",
    "link": "https://www.last.fm/music/Queen/_/My+Fairy+King"
}
```

### 3. База данных

Используется PostgreSQL. Структура базы данных создается автоматически при старте сервиса с помощью миграций. Таблица `songs` содержит следующие поля:
- `id` (primary key)
- `group` (название группы)
- `song` (название песни)
- `releaseDate` (дата выпуска)
- `text` (текст песни)
- `link` (ссылка на песню)

### 4. Логирование

Код покрыт debug- и info-логами. Пример логов:
```
[INFO] Starting to fetch songs
[DEBUG] Filters - Group: Muse, Song: Supermassive Black Hole, Offset: 0, Limit: 10
[INFO] Successfully fetched 1 songs
```

### 5. Конфигурация

Все настройки вынесены в `.env`-файл:
```env
DATABASE_URL=postgres://username:password@localhost:5432/music_library?sslmode=disable
LASTFM_API_KEY=YOUR_API_KEY_4444
```

### 6. Swagger-документация

Сгенерирована Swagger-документация для API. После запуска сервера документация доступна по адресу:
```
http://localhost:8080/swagger/index.html
```

---
