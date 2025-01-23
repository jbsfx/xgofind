# Project Name
PROJECT_NAME := xgofind

# Source Files
SRC := cmd/proxy/main.go

# Output Directory
BUILD_DIR := bin

# Default Binary Name
BINARY_WIN := $(BUILD_DIR)/$(PROJECT_NAME).exe
BINARY_LINUX := $(BUILD_DIR)/$(PROJECT_NAME)-linux
BINARY_MACOS := $(BUILD_DIR)/$(PROJECT_NAME)-macos

# Default Build
.PHONY: all clean build-windows build-linux build-macos

all: build-windows

# Build for Windows
build-windows:
	@echo "Building for Windows..."
	set GOOS=windows | set GOARCH=amd64 | go build -o $(BINARY_WIN) $(SRC)
	@echo "Binary created: $(BINARY_WIN)"

# Build for Linux
build-linux:
	@echo "Building for Linux..."
	set GOOS=linux && set GOARCH=amd64 && go build -o $(BINARY_LINUX) $(SRC)
	@echo "Binary created: $(BINARY_LINUX)"

# Build for macOS
build-macos:
	@echo "Building for macOS..."
	set GOOS=darwin && set GOARCH=amd64 && go build -o $(BINARY_MACOS) $(SRC)
	@echo "Binary created: $(BINARY_MACOS)"

# Clean the build output
clean:
	@echo "Cleaning up build directory..."
	del /Q $(BUILD_DIR)\*
	@echo "Cleaned."
