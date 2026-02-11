# 常用库演示 (libs-demo)

这个目录用于演示 Go 语言中处理**动态公式计算**的场景，使用了 `expr` 库。

它解决了 Go 语言原生不支持三元运算符 (`? :`) 的问题，非常适合用于规则引擎、价格计算、风控逻辑等场景。

## 依赖库

- **expr**: `github.com/antonmedv/expr` (高效的表达式求值引擎)

## 如何运行

由于这是一个新模块，你需要先下载依赖：

1.  初始化并下载依赖：
    ```bash
    go mod tidy
    ```

2.  运行代码：
    ```bash
    go run main.go
    ```

## 其他推荐库（未在此演示中包含）

针对你的技术栈，推荐以下官方/标准库：

- **Redis**: `github.com/redis/go-redis/v9`
- **InfluxDB**: `github.com/influxdata/influxdb-client-go/v2`
- **TDengine**: `github.com/taosdata/driver-go/v3`
- **MySQL**: `gorm.io/driver/mysql` (配合 GORM)
- **PostgreSQL**: `gorm.io/driver/postgres` (配合 GORM)
