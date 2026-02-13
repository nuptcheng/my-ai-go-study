# Go 语言项目开发规范与目录结构指南

## 1. 命名规范 (Naming Conventions)

Go 的命名哲学是：**简洁、清晰、统一**。

### 1.1 文件名
*   **规范**: 使用 **全小写**，单词间用 **下划线** (`_`) 分隔（Snake Case）。
*   **避免**: 驼峰式 (`modbusHandleService.go`) 或 中划线 (`modbus-handle-service.go`)。
*   **示例**:
    *   ✅ `modbus_handle_service.go`
    *   ✅ `user_controller.go`
    *   ❌ `modbusHandleService.go`
    *   ❌ `modbus-handle-service.go`
*   **特殊后缀**:
    *   测试文件必须以 `_test.go` 结尾。
    *   特定平台文件如 `file_windows.go`, `file_linux.go`。

### 1.2 常量 (Constants)
*   **规范**: 使用 **驼峰式命名 (CamelCase)**。
    *   **导出 (Public)**: 首字母大写 (`MaxLength`)。
    *   **非导出 (Private)**: 首字母小写 (`maxLength`)。
*   **避免**: 全大写下划线 (`MAX_LENGTH`)，除非是与 C 语言交互或极少数特定领域的惯例。
*   **示例**:
    ```go
    const MaxLength = 1024       // ✅ 导出的常量
    const defaultTimeout = 30    // ✅ 包内部使用的常量
    const MAX_BUFFER_SIZE = 2048 // ⚠️ 不推荐 (除非特殊原因)
    ```

### 1.3 变量 (Variables)
*   **规范**: 使用 **驼峰式命名 (CamelCase)**。
*   **缩写**: 常见缩写应保持一致的大小写（如 `url` 而不是 `Url`，`id` 而不是 `Id`）。
    *   ✅ `userID`, `xmlHTTPRequest`
    *   ❌ `userId`, `XmlHttpRequest`
*   **示例**:
    ```go
    var (
        userCount int
        APIKey    string // 缩写全大写
    )
    ```

### 1.4 结构体与接口 (Structs & Interfaces)
*   **规范**: 使用 **驼峰式命名 (CamelCase)**，首字母大写表示导出。
*   **接口**: 单方法接口通常以 `er` 结尾（如 `Reader`, `Writer`, `Formatter`）。
*   **示例**:
    ```go
    type UserService struct {}
    type Reader interface {
        Read(p []byte) (n int, err error)
    }
    ```

### 1.5 包名 (Package Names)
*   **规范**: **全小写**，**单数形式**，尽量简短。
*   **避免**: 下划线 (`my_package`) 或 驼峰 (`myPackage`)。
*   **示例**: `time`, `list`, `http`, `user`。

---

## 2. 项目目录结构规划 (Project Layout)

针对您的需求（配置、服务、工具、文档、AI、资源、发布），建议采用符合 Go 社区惯例（Standard Go Project Layout）的结构，同时兼顾您的特定业务模块。

### 推荐的目录树结构

```bash
my-project/
├── cmd/                        # 主程序入口 (Main Applications)
│   └── server/                 # 具体的应用名称
│       └── main.go             # 程序入口文件
├── internal/                   # 私有应用代码 (外部项目无法导入)
│   ├── config/                 # 配置结构定义与加载逻辑 (Config Structs)
│   ├── controller/             # API 接口层 (HTTP Handlers)
│   ├── service/                # 业务逻辑层 (Business Logic)
│   └── model/                  # 数据模型 (Data Models)
├── pkg/                        # 公共库代码 (可被外部项目导入)
│   ├── utils/                  # 通用工具 (建议按功能拆分，如 strutil, timeutil)
│   └── constants/              # 全局常量 (谨慎使用，避免循环依赖)
├── configs/                    # 配置文件模板 (JSON/YAML/TOML)
│   └── config.example.json
├── scripts/                    # 脚本文件 (对应您的 01-sh)
│   ├── docker-run.sh
│   └── build.sh
├── docs/                       # 文档 (对应您的 02-docs)
│   ├── api.md
│   └── architecture.md
├── ai/                         # AI 相关建议与模型 (对应您的 03-ai)
├── resources/                  # 静态资源/第三方文件 (对应您的 04-resources)
├── dist/                       # 构建产物 (对应您的 05-releases)
│   ├── my-app-linux-amd64
│   ├── my-app-linux-arm64
│   └── my-app-windows-amd64.exe
├── go.mod                      # 依赖管理文件
├── go.sum                      # 依赖版本锁定
└── README.md                   # 项目说明
```

### 目录映射与说明

| 您的构想 | 推荐名称 | 说明 |
| :--- | :--- | :--- |
| `config` (文件夹放json) | `configs/` | 用于存放 `.json`, `.yaml` 等静态配置文件。 |
| `config` (代码读取配置) | `internal/config/` | 存放配置的 Go 结构体定义和加载代码。 |
| `services` | `internal/service/` | 核心业务逻辑，通常不对外暴露，所以放在 `internal` 下。 |
| `controllers` | `internal/controller/` | 也就是 API 层，处理 HTTP 请求参数和响应。 |
| `utils` | `pkg/utils/` | 如果是通用的工具库，放在 `pkg` 下；如果是业务相关的，放在 `internal` 下。 |
| `constants` | `pkg/constants/` | 存放枚举或全局常量。 |
| `01-sh` | `scripts/` | Go 社区通常使用 `scripts` 目录存放 Shell/Python 脚本。 |
| `02-docs` | `docs/` | 标准的文档目录。 |
| `03-ai` | `ai/` | 保持原名即可，或者 `knowledge`。 |
| `04-resources` | `resources/` 或 `assets/` | 用于存放图片、模板等静态资源。 |
| `05-releases` | `dist/` 或 `bin/` | `dist` (distribution) 常用于存放发布包；`bin` 用于存放编译后的二进制文件。 |

### 关键原则
1.  **`internal` 目录**: 这是 Go 编译器强制的访问控制机制。放在 `internal` 目录下的代码，只有本项目内部可以 import，外部项目无法引用。这对于保护业务逻辑非常重要。
2.  **`cmd` 目录**: 将 `main.go` 放在 `cmd/应用名/` 下，可以支持一个项目编译出多个二进制程序（如 `cmd/server`, `cmd/worker`, `cmd/cli`）。
3.  **避免 `common` 或 `utils` 包**: 尽量按功能命名包（如 `date`, `file`, `security`），而不是把所有东西都塞进 `utils`，因为 `utils` 往往会变成垃圾堆且容易导致循环依赖。

---

## 3. 示例代码片段

### 常量定义 (pkg/constants/app.go)
```go
package constants

const (
    // AppName 应用名称 (导出)
    AppName = "ModbusService"
    // Version 版本号
    Version = "1.0.0"
)

// 状态枚举
type ServiceState int

const (
    StateUnknown ServiceState = iota
    StateRunning
    StateStopped
)
```

### 接口定义 (internal/service/modbus.go)
```go
package service

// ModbusHandler 定义接口
type ModbusHandler interface {
    ReadCoil(address uint16) (bool, error)
    WriteRegister(address uint16, value uint16) error
}
```
