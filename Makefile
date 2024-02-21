BINARY_NAME := nfexp

build:
	npm run build --prefix web
	go build -o $(BINARY_NAME) cmd/main.go

clean:
	go clean
	rm -f $(BINARY_NAME)
