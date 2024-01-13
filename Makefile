BINARY_NAME=phpenv
OUTPUT_DIR=build

build:
	GOOS=windows GOARCH=386 MODE=runtime go build -o ./${OUTPUT_DIR}/${BINARY_NAME}-386.exe ./cmd
	GOOS=windows GOARCH=amd64 go build -o ./${OUTPUT_DIR}/${BINARY_NAME}-amd64.exe ./cmd
	GOOS=darwin GOARCH=amd64 go build -o ./${OUTPUT_DIR}/${BINARY_NAME}-darwin ./cmd
	GOOS=linux GOARCH=386 go build -o ./${OUTPUT_DIR}/${BINARY_NAME}-linux-386 ./cmd
	GOOS=linux GOARCH=amd64 go build -o ./${OUTPUT_DIR}/${BINARY_NAME}-linux-amd64 ./cmd
	GOOS=linux GOARCH=arm go build -o ./${OUTPUT_DIR}/${BINARY_NAME}-linux-arm ./cmd
	GOOS=linux GOARCH=arm64 go build -o ./${OUTPUT_DIR}/${BINARY_NAME}-linux-arm64 ./cmd

.PHONY: build