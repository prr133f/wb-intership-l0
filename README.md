# Задание уровня L0
## Необходимые зависимости
- Docker
- NATS CLI
## Установка и запуск
Для начала, склонируйте репозиторий на локальную машину:
```sh
$ git clone git@github.com:prr133f/wb-intership-l0.git
```
Затем перейдите в директорию с проектом и создайте конфигуракионный файл окружения `secrets.env`:
```sh
$ cd wb-intership-l0
$ touch secrets.env
```
Необходимая структура файла выглядит следующим образом:
```
POSTGRES_USER=user
POSTGRES_PASSWORD=pass
POSTGRES_DB=db

NATS_USER=user
NATS_PASSWORD=pass

NATS_DSN=nats://${NATS_USER}:${NATS_PASSWORD}@nats:4222
POSTGRES_DSN=postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@db:5432/${POSTGRES_DB}
APP_PORT=3000
APP_STATUS=debug

GOOSE_DRIVER=postgres
GOOSE_DBSTRING=${POSTGRES_DSN}?sslmode=disable
```
Далее, запустите docker compose
```sh
$ docker compose up --build
```

После запуска контейнеров вы можете обращаться к:
- API по адресу `localhost:3000/v1/`
- NATS по адресу `0.0.0.0:4222`

## Бенчмарки
Бенчмарк проводился при помощи утилиты Vegeta и выдал следующие результаты

<details>
  <summary>График</summary>
  <img src="Vegeta Plot.png" alt="Benchmark plot" width="800">
</details>
