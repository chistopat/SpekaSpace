# CodeGeneration from scratch

---

### Цель - разобраться как работает кодогенерация

### As Is:
**Пустой проект, ветка master**

### To be:
**Запускаемый и тестируемый микросервис simple-gw**

---
Клонировать репозиторий
```
cd ~
git clone git@git.redmadrobot.com:internship/backend/lim-ext.git lim-ext
cd lim-ext
git checkout -b my_branch origin/codegen_from_scratch
```
---

Написать Makefile (смотри код в коммите)
**Важно** реализация очень простенькая, без аргумента `target=<service>`
можно что-нибудь сломать
---
Инициализировать структуру каталогов, сгенерировать файлы
make init нужен только при самом первом запуске проекта, но повторный вызов безопасный 
```
make init target=simple_gateway
make gen target=simple_gateway
```

---

Структура проекта в результате:

```
.
├── Makefile
├── README.md
├── doc
│   └── simple_gateway
│       └── openapi.yaml
├── entrypoint
│   └── simple_gateway
│       └── main.go
├── go.mod
├── go.sum
├── pkg
│   ├── http
│   │   └── server.go
│   ├── logger
│   │   └── logger.go
│   ├── openapi
│   │   └── openapi.go
│   ├── testing
│   │   └── testing.go
│   └── utils
│       └── utils.go
└── service
    └── simple_gateway
        ├── config
        │   └── config.go
        ├── generated
        │   ├── server.gen.go
        │   ├── spec.gen.go
        │   └── types.gen.go
        ├── handlers.go
        ├── service.go
        └── tests

```

---

```
go run ./entrypoint/simple_gateway/main.go
```


### ЗЫ
**эти файлы пишутся и изменяются вручную**
```
/cmd/simple_gateway/main.go
/service/simple_gateway/config/config.go
/pkg/*
/service/simple_gateway/service.go
/service/simple_gateway/handlers.go
```
