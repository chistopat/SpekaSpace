CODEGEN = github.com/deepmap/oapi-codegen/cmd/oapi-codegen
GENERATED_PATH = ./src/generated
APIDOC = ./api/openapi.yaml

.PHONY: clean
clean:
	if [ -d "$(GENERATED_PATH)" ]; then	rm -rf $(GENERATED_PATH); fi

generate:
	mkdir -p $(GENERATED_PATH)
	go run $(CODEGEN) -package generated -generate types $(APIDOC) > $(GENERATED_PATH)/types.gen.go
	go run $(CODEGEN) -package generated -generate chi-server $(APIDOC) > $(GENERATED_PATH)/server.gen.go
	go run $(CODEGEN) -package generated -generate spec $(APIDOC) > $(GENERATED_PATH)/spec.gen.go

generate_speka_space: clean generate
