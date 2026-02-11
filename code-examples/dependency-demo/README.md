# Go Dependency Management Demo

这个示例演示了 Go Modules 的基本用法。

## 核心文件

- `main.go`: 引入了外部库 `github.com/google/uuid`
- `go.mod`: 模块定义文件 (类似 package.json)
- `go.sum`: 依赖校验文件 (类似 package-lock.json, 运行 tidy 后生成)

## 如何运行

1. **初始化模块** (如果还没做):
   ```bash
   go mod init dependency-demo
   ```
   *(注: 示例中已包含初始化的 go.mod)*

2. **下载整理依赖**:
   这步最关键，它会分析代码中的 import，自动下载缺失的包并更新 go.mod/go.sum。
   ```bash
   go mod tidy
   ```

3. **运行代码**:
   ```bash
   go run main.go
   ```

## 观察结果

运行 `go mod tidy` 后，你会发现：
1. `go.mod` 文件中多了 `require github.com/google/uuid v1.x.x`
2. 目录下自动生成了 `go.sum` 文件
3. **注意**: 项目目录下**不会**出现类似 `node_modules` 的文件夹，源码被下载到了全局缓存 (`$GOPATH/pkg/mod`)。
