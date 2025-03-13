# Output binary name prefix
BIN=nnn

all: clean build wrapper updater

# Build for all platforms
build: build_mac_amd64 build_mac_arm64 build_linux_amd64 build_linux_arm64 build_windows_amd64 build_windows_arm64

build_mac_amd64:
	GOOS=darwin GOARCH=amd64 go build -o bin/$(BIN)_amd64_mac src/*.go

build_mac_arm64:
	GOOS=darwin GOARCH=arm64 go build -o bin/$(BIN)_arm64_mac src/*.go

build_linux_amd64:
	GOOS=linux GOARCH=amd64 go build -o bin/$(BIN)_amd64_linux src/*.go

build_linux_arm64:
	GOOS=linux GOARCH=arm64 go build -o bin/$(BIN)_arm64_linux src/*.go

build_windows_amd64:
	GOOS=windows GOARCH=amd64 go build -o bin/$(BIN)_amd64_win.exe src/*.go

build_windows_arm64:
	GOOS=windows GOARCH=arm64 go build -o bin/$(BIN)_arm64_win.exe src/*.go

# Build wrapper for all platforms
wrapper: wrapper_mac_amd64 wrapper_mac_arm64 wrapper_linux_amd64 wrapper_linux_arm64 wrapper_windows_amd64 wrapper_windows_arm64

wrapper_mac_amd64:
	GOOS=darwin GOARCH=amd64 go build -o bin/run_wrap_amd64_mac wrap/*.go

wrapper_mac_arm64:
	GOOS=darwin GOARCH=arm64 go build -o bin/run_wrap_arm64_mac wrap/*.go

wrapper_linux_amd64:
	GOOS=linux GOARCH=amd64 go build -o bin/run_wrap_amd64_linux wrap/*.go

wrapper_linux_arm64:
	GOOS=linux GOARCH=arm64 go build -o bin/run_wrap_arm64_linux wrap/*.go

wrapper_windows_amd64:
	GOOS=windows GOARCH=amd64 go build -o bin/run_wrap_amd64_win.exe wrap/*.go

wrapper_windows_arm64:
	GOOS=windows GOARCH=arm64 go build -o bin/run_wrap_arm64_win.exe wrap/*.go

# Build updater for all platforms
updater: updater_mac_amd64 updater_mac_arm64 updater_linux_amd64 updater_linux_arm64 updater_windows_amd64 updater_windows_arm64

updater_mac_amd64:
	GOOS=darwin GOARCH=amd64 go build -o bin/updater_amd64_mac updater/*.go

updater_mac_arm64:
	GOOS=darwin GOARCH=arm64 go build -o bin/updater_arm64_mac updater/*.go

updater_linux_amd64:
	GOOS=linux GOARCH=amd64 go build -o bin/updater_amd64_linux updater/*.go

updater_linux_arm64:
	GOOS=linux GOARCH=arm64 go build -o bin/updater_arm64_linux updater/*.go

updater_windows_amd64:
	GOOS=windows GOARCH=amd64 go build -o bin/updater_amd64_win.exe updater/*.go

updater_windows_arm64:
	GOOS=windows GOARCH=arm64 go build -o bin/updater_arm64_win.exe updater/*.go

# Clean binaries
clean:
	rm -rf bin/*

run:
	./bin/$(BIN)_amd64_mac run ls