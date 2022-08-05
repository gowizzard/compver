# Here you can build the binaries
build:
	GOOS=windows GOARCH=amd64 go build -o build/app-amd64.exe
	GOOS=linux GOARCH=amd64 go build -o build/compver-amd64-linux
	GOOS=darwin GOARCH=amd64 go build -o build/compver-amd64-darwin