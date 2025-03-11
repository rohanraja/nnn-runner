# Output binary name
BIN=nnn

all: build

# Build for macOS/Linux
build:
	GOOS=darwin GOARCH=arm64 go build -o $(BIN) *.go

# Clean binaries
clean:
	rm -f $(BIN)
