build:
	@echo "Building wasm"
	@GOARCH=wasm GOOS=js go build -v -o ./web/app.wasm .
	@echo "Building go app"
	@go build -o "portfolio"


run: build
	@echo "Running on local"
	./portfolio local


github: build
	@echo "Building for github"
	./portfolio github


clean:
	@go clean -v ./...
	-@rm ./portfolio
	@go mod tidy
