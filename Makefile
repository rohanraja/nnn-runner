# Output binary name
BIN=nnn

all: build

# Build for macOS/Linux
build:
	GOOS=darwin GOARCH=arm64 go build -o bin/$(BIN) src/*.go

wrapper:
	GOOS=darwin GOARCH=arm64 go build -o bin/run_wrap wrap/*.go

# Clean binaries
clean:
	rm -f $(BIN)

run:
	./bin/$(BIN) run ls