# Request 🛜

Simple HTTP client in Go (mini `curl` alternative).  
Простой HTTP-клиент на Go (аналог `curl`).

---

## Installation / Установка

```bash
git clone https://github.com/gox7/request.git
cd request
make build
````

Binary will be in `./build/request`.
Собранный бинарь будет в `./build/request`.

---

## Usage / Использование

```bash
./request [flags]
```

### Supported flags / Поддерживаемые флаги

| Flag            | Description (EN)              | Описание (RU)                      | Default  |
| --------------- | ----------------------------- | ---------------------------------- | -------- |
| `-X`            | HTTP method                   | HTTP-метод                         | `GET`    |
| `-H`            | Request header                | Заголовок запроса                  | `none`   |
| `-B`            | Request body                  | Тело запроса                       | `""`     |
| `-a`            | Address                       | Адрес                              | `none`   |
| `-p`            | Protocol (`http`)             | Протокол (`http`)                  | `http`   |
| `-t`            | Timeout (seconds)             | Таймаут (в секундах)               | `10`     |
| `-o`            | Export (filename or `stdout`) | Экспорт (файл или `stdout`)        | `stdout` |
| `-test_request` | Run built-in request test     | Запустить встроенный тест запроса  | `false`  |
| `-test_export`  | Run built-in export test      | Запустить встроенный тест экспорта | `false`  |

---

## Examples / Примеры

**GET request / GET-запрос**:

```bash
./request -X GET -a https://ifconfig.me/ip
```

**POST request with body / POST-запрос с телом**:

```bash
./request -X POST -a https://httpbin.org/post -B '{"hello":"world"}' -H "Content-Type:application/json"
```

**Save result to file / Сохранить результат в файл**:

```bash
./request -a https://ifconfig.me/ip -o result.txt
```

---

## Tests / Тесты

Run internal tests:
Запуск встроенных тестов:

```bash
./request -test_request
./request -test_export
```

---

## Project structure / Структура проекта

```
cmd/            # entrypoint / точка входа
internal/config # flags & config / парсинг флагов и конфиг
internal/request# request logic / логика HTTP-запроса
internal/exports# export results / вывод/сохранение результата
internal/util   # test utils / утилиты для тестов
tests/          # tests / тесты
```
