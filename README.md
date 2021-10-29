# LIM

Бэкэнд для приложения Lim.

## Architecture

Микросервисы на Go. Есть общий gateway, который не должен стать оркестратором, а вызывать в идеале 1 микросервис, который делает обход в глубину и собирает что нужно.

## Code flow rules

Именование коммитов - стандартное. Представляем, что сначала идет *When applied that commit must* Add new feature, начинаем с глагола в настоящем времени.

## OpenAPI

Хорошая практика писать спеку руками, а из нее генерить код. Попробуем делать так с помощью либы `oapi-codegen`. Тем не менее в версии 1.8 у нее сломана валидация API в случае использовании префикса (/api/v1), поэтому **юзаем 1.7.1 намеренно**

После изменений в `api/openapi/internals/*.yaml` нужно выполнить `make generate_mobile_gw` или `generate_admin_gw`.
Все файлы сгенерятся в `ms/mobile_gw(admin_gw)/generated`. Реализуем новые хендлеры в `handlers.go` (реулизуем интерфейс **ServerInterface**), больше никуда по идее лезть не нужно. Будет ли достаточно валидации из свагера большой вопрос, будем смотреть. Сгенерированные типы из `components` лежат в `ma/mobile_gw(admin_gw)/generated/types.gen.go`, рядом в `server.gen.go` соответственно интерфейс сервера.

Для методов в спеке нужно помечать - реализованы ли они сейчас, или только добавлены без реализации. Хорошим тоном было бы добавлять в методы без реализации мок с данными из example.

В конце изменений поднимаемся в начало файла и вносим новую версию в CHANGELOG. По нему CI потом делаем оповещение.

## Generation proto
Настройки генерации берутся из файла buf.gen.yaml. Генерируется go и go-grpc в `api/proto/ms/<service_name>/`
```bash
buf generate
```

## Run linters
```bash
golangci-lint run
```

## Generate client from Openapi
```bash
make generate_marketplace
```

Пример использования
```go
import 	eyebuyClient "git.redmadrobot.com/internship/backend/lim/ms/authenticator/service/authenticator/client"

marketplaceConfig := marketplaceClient.Configuration{
    BasePath: fmt.Sprintf("%s://%s:%s/%s", cfg.Server.Scheme, cfg.Server.Host, cfg.Server.Port, cfg.Server.URL),
}

client = marketplaceClient.NewAPIClient(&marketplaceConfig)
_, _, err := s.client.MarketplaceApi.ListProducts(ctx)
```

## Generate ORM
Если вы поменяли или добавили схему данных в `ms/<service>/schema/, вам нужно ее сгенерировать заново. В качестве значения параметра `SERVICE_NAME` указать имя сервиса.
```bash
make generate_orm SERVICE_NAME=authenticator
```

## Tests
Для работы с функциональными тестами используется библиотека `https://github.com/lamoda/gonkey`.
Поднимается два сервиса для работы со streams и marketplace. Proto клиенты передаются в gateway сервис.
Для эмуляции 3th party api сервисов используется mock от `gonkey`, сервис и ответ указывается в самом case.

Запуск тестов локально должен производиться с заимпортированными переменными из .env файла. Также нужно поднять из docker-compose сервис с базой `test-lim-postgres`. Здесь 2 варианта.
1. Поднимаем все сервисы локально, и 2 базы. на 32м порту крутится база для локальных сервисов, на 33м - для тестов. Тогда в настройках энвов нужно заменить все порты коннекта к базам на 33.
2. Поднимаем только тестовую базу на 5432 порту, и все тесты должны заработать.

## Эмуляция marketplace
Установить https://github.com/stoplightio/prism
```bash
npm install -g @stoplight/prism-cli
```

Запуск эмуляции:
```bash
prism mock lim/api/openapi/externals/marketplace.yaml
```
