BINARY_NAME := nfexp

build:
	npm run build --prefix web
	go build -o $(BINARY_NAME) cmd/main.go

buildall: build
	GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME)-linux-amd64 cmd/main.go
	GOOS=darwin GOARCH=arm64 go build -o $(BINARY_NAME)-darwin-arm64 cmd/main.go
	GOOS=windows GOARCH=amd64 go build -o $(BINARY_NAME)-windows-amd64.exe cmd/main.go

clean:
	go clean
	rm -f $(BINARY_NAME)
