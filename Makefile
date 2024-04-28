run: build
	@./bin/app

build:
	@go build -o bin/app ./cmd/api/main.go

test:
	@go clean -testcache
	@go test ./... -cover

cover: 
	@go clean -testcache
	@go test ./... -coverprofile=coverage.out
	@go tool cover -html=coverage.out

