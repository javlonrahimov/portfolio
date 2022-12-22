build:
	@GOARCH=wasm GOOS=js go build -v -o web/app.wasm ./src
	@echo "\033[94m• Building go-app documentation\033[00m"
	@go build -o portfolio ./src

run: build
	@echo "\033[94m• Running go-app portfolio server\033[00m"
	@./portfolio local

github: build
	@echo "\033[94m• Generating GitHub Pages\033[00m"
	@./portfolio github

clean:
	@go clean -v ./...
	-@rm portfolio
	@go mod tidy
