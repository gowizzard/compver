# Here you can reformat, check or build the binary
BINARY_NAME=compver

fmt:
	go fmt ./...

vet:
	go vet ./...

lint:
	golangci-lint run ./...

test-build:
	go build -o /usr/local/bin/${BINARY_NAME}-testing
	chmod +x /usr/local/bin/${BINARY_NAME}-testing

test-remove:
	rm /usr/local/bin/gogs-testing

build:
	GOOS=windows GOARCH=amd64 go build -o build/${BINARY_NAME}-amd64-windows.exe
	GOOS=linux GOARCH=amd64 go build -o build/${BINARY_NAME}-amd64-linux
	GOOS=darwin GOARCH=amd64 go build -o build/${BINARY_NAME}-amd64-darwin