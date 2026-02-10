# Makefile 示例：演示 Go 的交叉编译

# 项目名称
APP_NAME := myapp
# 源码入口
MAIN_SRC := ./cmd/myapp/main.go
# 输出目录
BUILD_DIR := ./build

.PHONY: all clean build-linux-amd64 build-linux-arm64 build-linux-arm

all: clean build-linux-amd64 build-linux-arm64 build-linux-arm

# 1. 编译 Linux x86_64 (Ubuntu 20.04 服务器)
build-linux-amd64:
	@echo "Building for Linux amd64..."
	@mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(APP_NAME)-linux-amd64 $(MAIN_SRC)
	@echo "Done: $(BUILD_DIR)/$(APP_NAME)-linux-amd64"

# 2. 编译 Linux ARM64 (ARM64 网关)
build-linux-arm64:
	@echo "Building for Linux arm64..."
	@mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=arm64 go build -o $(BUILD_DIR)/$(APP_NAME)-linux-arm64 $(MAIN_SRC)
	@echo "Done: $(BUILD_DIR)/$(APP_NAME)-linux-arm64"

# 3. 编译 Linux ARM32 (ARM32 网关, 通常是 GOARCH=arm GOARM=7 或 6)
# 这里假设是 ARMv7
build-linux-arm:
	@echo "Building for Linux arm (v7)..."
	@mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=arm GOARM=7 go build -o $(BUILD_DIR)/$(APP_NAME)-linux-arm $(MAIN_SRC)
	@echo "Done: $(BUILD_DIR)/$(APP_NAME)-linux-arm"

clean:
	@rm -rf $(BUILD_DIR)
	@echo "Cleaned build directory."

test:
	go test ./...
