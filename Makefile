SERVICE = $(target)
CODEGEN = github.com/deepmap/oapi-codegen/cmd/oapi-codegen
SERVICE_PATH = ./services/$(SERVICE)
GENERATED_PATH = $(SERVICE_PATH)/generated
DOCS_PATH = ./docs/$(SERVICE)
APIDOC = ./docs/$(SERVICE)/openapi.yaml
TESTS_PATH = $(SERVICE_PATH)/tests

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
	mkdir -p $(DOCS_PATH)
	mkdir -p $(SERVICE_PATH)
	mkdir -p $(TESTS_PATH)
	if [ ! -f "$(DOCS_PATH)/openapi.yaml" ]; then touch $(DOCS_PATH)/openapi.yaml; fi
	if [ ! -f "$(SERVICE_PATH)" ]; then touch $(SERVICE_PATH)/main.go; fi
	if [ ! -f "$(SERVICE_PATH)" ]; then touch $(SERVICE_PATH)/config.go; fi
	if [ ! -f "$(SERVICE_PATH)" ]; then touch $(SERVICE_PATH)/handlers.go; fi
	git add .
