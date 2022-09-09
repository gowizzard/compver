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
	rm /usr/local/bin/${BINARY_NAME}-testing

build:
	go build -o ${BINARY_NAME}

publish:
	GOOS=windows GOARCH=amd64 go build -o build/${BINARY_NAME}-amd64-windows.exe
	GOOS=linux GOARCH=amd64 go build -o build/${BINARY_NAME}-amd64-linux
	GOOS=darwin GOARCH=amd64 go build -o build/${BINARY_NAME}-amd64-darwin

docker-test-build:
	docker build -t ${BINARY_NAME}-testing .

docker-test-run:
	docker run --name ${BINARY_NAME}-testing1 -e GITHUB_ACTIONS=true ${BINARY_NAME}-testing:latest -compare -version1 v1.2.0 -version2 v1.2.5
	docker run --name ${BINARY_NAME}-testing2 -e GITHUB_ACTIONS=true ${BINARY_NAME}-testing:latest -core -version1 v1.11.0 -block minor

docker-test-remove:
	docker stop ${BINARY_NAME}-testing1 ${BINARY_NAME}-testing2
	docker rm ${BINARY_NAME}-testing1 ${BINARY_NAME}-testing2
