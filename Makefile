build:
	@go build -o bin/ggcommerce

run: build
	@./bin/ggcommerce

test:
	@go test -v ./...
