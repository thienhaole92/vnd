.PHONY: build,run-local,pre-commit

SERVICE_NAME = rest

# run
build:
	env GOOS=$(TARGETOS) GOARCH=$(TARGETARCH) go build -o $(SERVICE_NAME)

run-local:
	$(shell export $$(grep -v '#' .env.local | xargs -I % echo -n '% ') && ./$(SERVICE_NAME))

pre-commit:
	go mod tidy
	go vet ./...
	go fmt ./...

# doc
pre-doc-go:
	go install github.com/swaggo/swag/cmd/swag@v1.16.2

generate-doc:
	swag init -d ./,./internal/service/ -q -o api-spec -ot yaml --parseDependency --parseInternal
