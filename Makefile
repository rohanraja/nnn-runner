# Output binary name prefix
BIN=nn

all: clean build wrapper updater

# Build for all platforms
build: build_mac_amd64 build_mac_arm64 build_linux_amd64 build_linux_arm64 build_windows_amd64 build_windows_arm64

# Build wrapper for all platforms
wrapper: wrapper_mac_amd64 wrapper_mac_arm64 wrapper_linux_amd64 wrapper_linux_arm64 wrapper_windows_amd64 wrapper_windows_arm64

# Build updater for all platforms
updater: updater_mac_amd64 updater_mac_arm64 updater_linux_amd64 updater_linux_arm64 updater_windows_amd64 updater_windows_arm64

# New platform-specific targets that build all components for that platform
mac_amd64: build_mac_amd64 wrapper_mac_amd64 updater_mac_amd64
mac_arm64: build_mac_arm64 wrapper_mac_arm64 updater_mac_arm64
linux_amd64: build_linux_amd64 wrapper_linux_amd64 updater_linux_amd64
linux_arm64: build_linux_arm64 wrapper_linux_arm64 updater_linux_arm64
windows_amd64: build_windows_amd64 wrapper_windows_amd64 updater_windows_amd64
windows_arm64: build_windows_arm64 wrapper_windows_arm64 updater_windows_arm64

# New platform-grouped all target
all_platforms: mac_amd64 mac_arm64 linux_amd64 linux_arm64 windows_amd64 windows_arm64

# Individual component targets
build_mac_amd64:
	mkdir -p bin/amd64/darwin
	GOOS=darwin GOARCH=amd64 go build -o bin/amd64/darwin/$(BIN) src/*.go

build_mac_arm64:
	mkdir -p bin/arm64/darwin
	GOOS=darwin GOARCH=arm64 go build -o bin/arm64/darwin/$(BIN) src/*.go

build_linux_amd64:
	mkdir -p bin/amd64/linux
	GOOS=linux GOARCH=amd64 go build -o bin/amd64/linux/$(BIN) src/*.go

build_linux_arm64:
	mkdir -p bin/arm64/linux
	GOOS=linux GOARCH=arm64 go build -o bin/arm64/linux/$(BIN) src/*.go

build_windows_amd64:
	mkdir -p bin/amd64/windows
	GOOS=windows GOARCH=amd64 go build -o bin/amd64/windows/$(BIN).exe src/*.go

build_windows_arm64:
	mkdir -p bin/arm64/windows
	GOOS=windows GOARCH=arm64 go build -o bin/arm64/windows/$(BIN).exe src/*.go

wrapper_mac_amd64:
	mkdir -p bin/amd64/darwin
	GOOS=darwin GOARCH=amd64 go build -o bin/amd64/darwin/run_wrap wrap/*.go

wrapper_mac_arm64:
	mkdir -p bin/arm64/darwin
	GOOS=darwin GOARCH=arm64 go build -o bin/arm64/darwin/run_wrap wrap/*.go

wrapper_linux_amd64:
	mkdir -p bin/amd64/linux
	GOOS=linux GOARCH=amd64 go build -o bin/amd64/linux/run_wrap wrap/*.go

wrapper_linux_arm64:
	mkdir -p bin/arm64/linux
	GOOS=linux GOARCH=arm64 go build -o bin/arm64/linux/run_wrap wrap/*.go

wrapper_windows_amd64:
	mkdir -p bin/amd64/windows
	GOOS=windows GOARCH=amd64 go build -o bin/amd64/windows/run_wrap.exe wrap/*.go

wrapper_windows_arm64:
	mkdir -p bin/arm64/windows
	GOOS=windows GOARCH=arm64 go build -o bin/arm64/windows/run_wrap.exe wrap/*.go

updater_mac_amd64:
	mkdir -p bin/amd64/darwin
	GOOS=darwin GOARCH=amd64 go build -o bin/amd64/darwin/updater updater/*.go

updater_mac_arm64:
	mkdir -p bin/arm64/darwin
	GOOS=darwin GOARCH=arm64 go build -o bin/arm64/darwin/updater updater/*.go

updater_linux_amd64:
	mkdir -p bin/amd64/linux
	GOOS=linux GOARCH=amd64 go build -o bin/amd64/linux/updater updater/*.go

updater_linux_arm64:
	mkdir -p bin/arm64/linux
	GOOS=linux GOARCH=arm64 go build -o bin/arm64/linux/updater updater/*.go

updater_windows_amd64:
	mkdir -p bin/amd64/windows
	GOOS=windows GOARCH=amd64 go build -o bin/amd64/windows/updater.exe updater/*.go

updater_windows_arm64:
	mkdir -p bin/arm64/windows
	GOOS=windows GOARCH=arm64 go build -o bin/arm64/windows/updater.exe updater/*.go

# Clean binaries
clean:
	rm -rf bin/*

run:
	./bin/amd64/darwin/$(BIN) run ls