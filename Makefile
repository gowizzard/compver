# Here you can reformat, check or build the binary
BINARY_NAME=compver
GIT_TAG=$(shell git describe --tags --abbrev=0)
LDFLAGS=-ldflags "-X 'github.com/gowizzard/${BINARY_NAME}/v3/build_information.Version=${GIT_TAG}'"

fmt:
	go fmt ./...

vet:
	go vet ./...

test:
	go test ./...

lint:
	golangci-lint run ./...

version:
	@echo "version: ${GIT_TAG}"

test-build:
	go build ${LDFLAGS} -o /usr/local/bin/${BINARY_NAME}-testing
	chmod +x /usr/local/bin/${BINARY_NAME}-testing

test-run:
	${BINARY_NAME}-testing -version
	${BINARY_NAME}-testing -compare -version1 v1.2.0 -version2 v1.2.5
	${BINARY_NAME}-testing -core -block minor -version1 v1.11.0

test-remove:
	rm /usr/local/bin/${BINARY_NAME}-testing

test-all: test-build test-run test-remove

build:
	go build ${LDFLAGS} -o ${BINARY_NAME}

cross-compile:
	GOOS=windows GOARCH=amd64 go build ${LDFLAGS} -o build/cross-compile/${BINARY_NAME}-amd64-windows.exe
	GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o build/cross-compile/${BINARY_NAME}-amd64-linux
	GOOS=darwin GOARCH=amd64 go build ${LDFLAGS} -o build/cross-compile/${BINARY_NAME}-amd64-darwin

docker-test-build:
	docker build -t ${BINARY_NAME}-testing .

docker-test-run:
	docker run --name ${BINARY_NAME}-testing1 -e GITHUB_ACTIONS=true ${BINARY_NAME}-testing:latest -compare -version1 v1.2.0 -version2 v1.2.5
	docker run --name ${BINARY_NAME}-testing2 -e GITHUB_ACTIONS=true ${BINARY_NAME}-testing:latest -core -block minor -version1 v1.11.0

docker-test-stop:
	docker stop ${BINARY_NAME}-testing1 ${BINARY_NAME}-testing2

docker-test-remove:
	docker rm ${BINARY_NAME}-testing1 ${BINARY_NAME}-testing2

docker-test-all: docker-test-build docker-test-run docker-test-stop docker-test-remove