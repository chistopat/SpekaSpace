# CodeGeneration from scratch

---

### Цель - разобраться как работает кодогенерация

### As Is:
**Пустой проект, ветка master**

### To be:
**Запускаемый и тестируемый микроесервис simple-gw**

---
Клонируем репозиторий
```
cd ~
git clone git@git.redmadrobot.com:internship/backend/lim-ext.git lim-ext
cd lim-ext
git chekout -b codegen_from_scratch
```
---

Инициализация golang
```
go mod init git.redmadrobot.com/internship/backend/lim-ext
go mod tidy
```

Создадим простую структуру директорий:
```
.
├── ReadMe.md
├── docs
│   └── simple_gateway
│       └── openapi.yaml
├── go.mod
└── services
    └── simple_gateway
```

В директории `docs` живет документация для api

В `services` скоро появится код нашего сервиса

`simple_gateway` - название нашего сервиса, оно же будет являться его идентификатором

---
Чтобы случилось чудо, нужно установить пакет **oapi-codegen**, указать ему где взять документацию и куда положить сгенерированный код
```
go get github.com/deepmap/oapi-codegen
```

Команды довольно длинные, поэтому для удобства создадим Makefile, там будут содержаться инструкции по работе с нашим проектом

```
make gen
```

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
        └── generated
            ├── server.gen.go
            ├── spec.gen.go
            └── types.gen.go
```

It works, в каталоге сервиса появилась директория generated
добавим эти файлы в .gitignore, чтобы они не мешались на код-ревью
и не мозолили глаза в gitlab, незачем людям тратить время на чтение кода
предназначенного для машины


