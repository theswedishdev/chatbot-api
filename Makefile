BINARY_NAME=chatbot-api

build: build_darwin build_linux

build_darwin:
	@echo "Building Darwin"
	env GOOS=darwin go build -o bin/${BINARY_NAME}_darwin

build_linux:
	@echo "Building Linux"
	env GOOS=linux go build -o bin/${BINARY_NAME}_linux

clean:
	go clean
	rm -r bin

deps:
	go get .

run:
	go run .
