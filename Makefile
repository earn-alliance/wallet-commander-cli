BINARY_NAME=wallet-commander
SRC_PATH=cmd/wallet_commander

build:
	GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME}-darwin ${SRC_PATH}/main.go
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}-linux ${SRC_PATH}/main.go
	GOARCH=amd64 GOOS=windows go build -o ${BINARY_NAME}-windows ${SRC_PATH}/main.go

run:
	./${BINARY_NAME}

build_and_run: build run

clean:
	go clean
	rm ${BINARY_NAME}-darwin
	rm ${BINARY_NAME}-linux
	rm ${BINARY_NAME}-windows

test:
	go test ./...

test_coverage:
	go test ./... -coverprofile=coverage.out

dep:
	go mod download

vet:
	go vet

lint:
	golangci-lint run --enable-all