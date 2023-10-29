BINARY_NAME=goperson

build:
	GOARCH=amd64 GOOS=darwin go build -o ./bin/${BINARY_NAME}-darwin ./cmd/go/
	GOARCH=amd64 GOOS=linux go build -o ./bin/${BINARY_NAME}-linux ./cmd/go/ 
	GOACH=amd64 GOOS=windows go build -o ./bin/${BINARY_NAME}-windows ./cmd/go/


run: build
	./bin/${BINARY_NAME}-linux

clean:
	go clean
	rm ./bin/${BINARY_NAME}-darwin
	rm ./bin/${BINARY_NAME}-linux
	rm ./bin/${BINARY_NAME}-windows

dep:s
	go mod download

vet:
	go vet

lint:
	golangci-lint run --enable-all