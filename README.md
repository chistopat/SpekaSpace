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
git chekout -b codegen_from_scratch
```
---

Инициализировать приложение golang
```
go mod init git.redmadrobot.com/internship/backend/lim-ext
go mod tidy
```
---

Написать Makefile (смотри код в коммите)
**Важно** реализация очень простенькая, без аргумента `target=<service>`
можно что-нибудь сломать
---
Инициализировать структуру каталогов, сгенерировать файлы
```
make init target=simple_gateway
make gen target=simple_gateway
```

---

Структура проекта в результате:

```
.
├── Makefile
├── ReadMe.md
├── docs
│   └── simple_gateway
│       └── openapi.yaml
├── go.mod
├── go.sum
└── services
    └── simple_gateway
        ├── config.go
        ├── generated
        │   ├── server.gen.go
        │   ├── spec.gen.go
        │   └── types.gen.go
        ├── handlers.go
        ├── main.go
        └── tests

```

---

```
go run ./entrypoint/simple_gateway/main.go
```