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

Создадим простую структуру директорий:
```
.
├── ReadMe.md
├── docs
│   └── simple_gateway
│       └── openapi.yaml
└── services
    └── simple_gateway
```

В директории `docs` живет документация для api

В `services` скоро появится код нашего сервиса

`simple_gateway` - название нашего сервиса, оно же будет являться его идентификатором

---
Чтобы случилось чудо, нужно установить пакет **oapi-codegen**, указать ему где взять документацию и куда положить сгенерированный код
Команды довольно длинные, поэтому для удобства создадим Makefile, там будут содержаться инструкции по работе с нашим проектом


