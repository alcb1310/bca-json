run: build
	./build/main

build: clean
	@go build -o build/main src/api/main.go

clean:
	@rm -rf build

watch:
	@air
