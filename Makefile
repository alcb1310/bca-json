mocks: clean
	@mockery --with-expecter=true --all

run: build
	./build/main

build: clean
	@go build -o build/main src/api/main.go

clean:
	@rm -rf build mocks

watch:
	@air

unit-test:
	@go clean -testcache
	@go test `go list ./... | grep -v ./src/api | grep -v ./internals/database | grep -v ./mocks | grep -v ./tests | grep -v ./externals`

integration-test:
	@go clean -testcache
	@go test ./tests/integration/...
