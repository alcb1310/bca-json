run: build
	./build/main

build: clean
	@go build -o build/main src/api/main.go

clean:
	@rm -rf build

watch:
	@air

unit-test:
	@go clean -testcache
	@go test `go list ./... | grep -v ./src/api | grep -v ./internals/database | grep -v ./mocks | grep -v ./tests | grep -v ./externals`

mocks:
	@mockery --with-expecter=true --all
