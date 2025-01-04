GO=go
BUILD_FLAGS=-ldflags="-s -w"
OUTPUT_FILE=hook007

all: clean build

build: build-linux-amd64 build-linux-arm64  build-darwin-amd64

# linux
build-linux-amd64:
	GOOS=linux GOARCH=amd64 $(GO) build $(BUILD_FLAGS) -o ./release/$(OUTPUT_FILE)-linux-amd64

build-linux-arm64:
	GOOS=linux GOARCH=arm64 $(GO) build $(BUILD_FLAGS) -o ./release/$(OUTPUT_FILE)-linux-arm64

# darwin
build-darwin-amd64:
	GOOS=darwin GOARCH=amd64 $(GO) build $(BUILD_FLAGS) -o ./release/$(OUTPUT_FILE)-darwin-amd64

clean:
	go clean
	rm -rf ./release
	mkdir -p ./release
