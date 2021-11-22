SERVICE = $(target)
CODEGEN = github.com/deepmap/oapi-codegen/cmd/oapi-codegen
SERVICE_PATH = ./service/$(SERVICE)
GENERATED_PATH = $(SERVICE_PATH)/generated
DOC_PATH = ./doc/$(SERVICE)
APIDOC = $(DOC_PATH)/openapi.yaml
TESTS_PATH = $(SERVICE_PATH)/tests
ENTRYPOINT_PATH = ./entrypoint/$(SERVICE)
CONFIG_PATH = $(SERVICE_PATH)/config
CONFIG_FILE = $(CONFIG_PATH)/config.go
PKG_PATH = ./pkg

.PHONY: clean
clean:
	if [ -d "$(GENERATED_PATH)" ]; then	rm -rf $(GENERATED_PATH); fi

generate:
	mkdir -p $(GENERATED_PATH)
	go run $(CODEGEN) -package generated -generate types $(APIDOC) > $(GENERATED_PATH)/types.gen.go
	go run $(CODEGEN) -package generated -generate chi-server $(APIDOC) > $(GENERATED_PATH)/server.gen.go
	go run $(CODEGEN) -package generated -generate spec $(APIDOC) > $(GENERATED_PATH)/spec.gen.go

.PHONY: gen
gen: clean generate

.PHONY: gen
init:
	mkdir -p $(DOC_PATH)
	mkdir -p $(SERVICE_PATH)
	mkdir -p $(TESTS_PATH)
	mkdir -p $(ENTRYPOINT_PATH)
	mkdir -p $(PKG_PATH)
	mkdir -p $(CONFIG_PATH)
	if [ ! -f "$(APIDOC)" ]; then touch $(APIDOC); fi
	if [ ! -f "$(CONFIG_FILE)" ]; then touch $(CONFIG_FILE); fi
	if [ ! -f "$(SERVICE_PATH)" ]; then touch $(SERVICE_PATH)/handlers.go; fi
	if [ ! -f "$(SERVICE_PATH)" ]; then touch $(SERVICE_PATH)/service.go; fi
	if [ ! -f "$(ENTRYPOINT_PATH)/main.go" ]; then touch $(ENTRYPOINT_PATH)/main.go; fi

	git add .
