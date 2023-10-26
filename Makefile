BINARY_NAME=goperson

build:
	go build -o ./bin/${BINARY_NAME}


run: build
	./bin/${BINARY_NAME}

clean:
	go clean
	rm ./bin/${BINARY_NAME}

dep:
	go mod download

vet:
	go vet

lint:
	golangci-lint run --enable-all