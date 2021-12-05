include .env

CODEGEN = github.com/deepmap/oapi-codegen/cmd/oapi-codegen
GENERATED_PATH = ./src/generated
ENT_GENERATED_PATH = ./src/repository/ent
APIDOC = ./api/openapi.yaml
PID = ./.pid
MAIN = ./cmd/main.go
GOBIN = ./bin/main

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

_test:
	go run github.com/lamoda/gonkey -host ${SERVER_HOST}:${SERVER_PORT} -tests tests/cases -env-file .env

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
