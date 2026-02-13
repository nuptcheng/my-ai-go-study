/*
 * @Author: BecomeBamboo
 * @Date: 2026-02-13 11:02:37
 * @LastEditTime: 2026-02-13 11:08:56
 * @Description: 
 */
package formulas // 👈 包名 (Package Name)

// 关于包名与文件名的关系：
// 1. 包名 (package formulas): 决定了其他代码如何引用它。这里是 formulas.InitFromRedis()。
//    通常建议包名和目录名 (formulas) 保持一致。
// 2. 文件名 (store.go): 只是这个文件的名字。
//    同一个包 (目录) 下可以有多个文件 (如 logic.go, types.go)，它们都属于 package formulas。
//    文件名不影响调用方式。

import (
	"encoding/json" // 👈 这就是 Go 标准库内置的 JSON 解析包
	"fmt"
	"log"
	"time"
)

// FormulaDto 定义计算公式结构体
// 命名偏好：为了明确这是一个数据对象 (Data Object)，我们加上 Dto 后缀。
// 这在团队开发中完全没问题，只要保持统一即可。
type FormulaDto struct {
	FormulaStr string `json:"formulaStr"` 
	ResultStr  string `json:"resultStr"`  
	Cron       string `json:"cron"`       
}

// GlobalCache 是一个包级变量
var GlobalCache []FormulaDto

// InitFromRedis 模拟从 Redis 加载数据
func InitFromRedis() {
	fmt.Println("[Config] 正在连接 Redis 加载公式配置...")
	time.Sleep(500 * time.Millisecond)

	// 1. 模拟从 Redis 获取到的原始 JSON 字符串
	// 这是一个 JSON 数组字符串
	mockRedisData := `[
		{
			"formulaStr": "a + b",
			"resultStr": "point_001",
			"cron": "*/5 * * * *"
		},
		{
			"formulaStr": "temperature * 1.8 + 32",
			"resultStr": "point_002",
			"cron": "0 0 * * *"
		}
	]`

	// 2. 使用 encoding/json 包进行解析 (Unmarshal)
	// Unmarshal 的作用：把 JSON 字符串 (字节切片) -> 转换成 Go 的结构体对象
	var data []FormulaDto
	err := json.Unmarshal([]byte(mockRedisData), &data)
	if err != nil {
		log.Fatalf("JSON 解析失败: %v", err)
	}

	GlobalCache = data
	fmt.Printf("[Config] 加载完成，共 %d 条公式\n", len(GlobalCache))
}

// GetByIndex 封装一个获取方法
func GetByIndex(index int) *FormulaDto {
	if index >= 0 && index < len(GlobalCache) {
		return &GlobalCache[index]
	}
	return nil
}
