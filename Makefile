# Here you can reformat, check or build the binary.
BINARY_NAME=compver
LIST=$(shell go list)
GIT_TAG=$(shell git describe --tags --abbrev=0)
VERSION=$(if $(GIT_TAG),$(GIT_TAG),unavailable)
LDFLAGS=-ldflags "-X '${LIST}/build_information.Version=${VERSION}'"
DOCKER_HUB_USERNAME=gowizzard

fmt:
	@go fmt ./...

vet:
	@go vet ./...

test:
	go test -v ./... -bench=.

coverage:
	@go test ./... -coverprofile=coverage.out
	@go tool cover -html=coverage.out

doc:
	@godoc -play=true -goroot=/usr/local/go -http=:6060

version:
	@echo "version: ${VERSION}"

test-build:
	go build ${LDFLAGS} -o /usr/local/bin/${BINARY_NAME}-testing
	chmod +x /usr/local/bin/${BINARY_NAME}-testing

test-run:
	${BINARY_NAME}-testing -version
	${BINARY_NAME}-testing -compare -version1 v1.2.0 -version2 v1.2.5 -trim -prefix v
	${BINARY_NAME}-testing -core -block minor -version1 1.11.0

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
	docker run --name ${BINARY_NAME}-testing1 ${BINARY_NAME}-testing:latest -compare -version1 1.2.0 -version2 1.2.5
	docker run --name ${BINARY_NAME}-testing2 ${BINARY_NAME}-testing:latest -core -block minor -version1 v1.11.0 -trim -prefix v

docker-test-stop:
	docker stop ${BINARY_NAME}-testing1 ${BINARY_NAME}-testing2

docker-test-remove:
	docker rm ${BINARY_NAME}-testing1 ${BINARY_NAME}-testing2

docker-test-all: docker-test-build docker-test-run docker-test-stop docker-test-remove

docker-build:
	docker build -t ${DOCKER_HUB_USERNAME}/${BINARY_NAME}:${VERSION} -t ${DOCKER_HUB_USERNAME}/${BINARY_NAME}:latest .

docker-push:
	docker push -a ${DOCKER_HUB_USERNAME}/${BINARY_NAME}