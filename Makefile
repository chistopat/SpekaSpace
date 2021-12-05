include .env

CODEGEN = github.com/deepmap/oapi-codegen/cmd/oapi-codegen
GENERATED_PATH = ./src/generated
ENT_GENERATED_PATH = ./src/repository/ent
APIDOC = ./api/openapi.yaml
PID = ./.pid
MAIN = ./cmd/main.go
GOBIN = ./bin/main
DSN2="host=${DB_HOST} port=${DB_PORT} dbname=${DB_NAME} sslmode=${DB_SSL_MODE} user=${DB_USER} password=${DB_PASSWORD}"
DSN = postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}

.PHONY: clean
clean:
	if [ -d "$(GENERATED_PATH)" ]; then	rm -rf $(GENERATED_PATH); fi
	if [ -d "$(ENT_GENERATED_PATH)" ]; then	rm -rf $(ENT_GENERATED_PATH); fi


generate:
	mkdir -p $(GENERATED_PATH)
	go run $(CODEGEN) -package generated -generate types $(APIDOC) > $(GENERATED_PATH)/types.gen.go
	go run $(CODEGEN) -package generated -generate chi-server $(APIDOC) > $(GENERATED_PATH)/server.gen.go
	go run $(CODEGEN) -package generated -generate spec $(APIDOC) > $(GENERATED_PATH)/spec.gen.go
	go run -mod=mod entgo.io/ent/cmd/ent generate ./src/repository/schema --target=$(ENT_GENERATED_PATH)

gen: clean generate

GONKEY_DB_PARAMS = -env-file .env -db_dsn $(DSN2) -fixtures ./tests/fixtures

_test:
	go run gonkey/main.go -v -debug -host ${SERVER_HOST}:${SERVER_PORT} -tests tests/cases $(GONKEY_DB_PARAMS)

run:
	go run $(MAIN)

build:
	go build -o $(GOBIN) $(MAIN)

start:
	@$(GOBIN) 2>&1 & echo $$! > $(PID)
	@cat $(PID) | sed "/^/s/^/  \>  PID: /"

stop:
	@-touch $(PID)
	@-kill `cat $(PID)` || true
	@-rm $(PID)

restart: stop build start

test: rebuild _test

set:
	set -a && source .env && set +a

rebuild: stop set gen build start

export:
	export $(grep -v '^#' .env | xargs -0)
