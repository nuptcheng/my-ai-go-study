/*
 * @Author: BecomeBamboo
 * @Date: 2026-02-12 16:00:23
 * @LastEditTime: 2026-02-12 16:03:34
 * @Description: 
 */
package main

import (
	"encoding/json"
	"fmt"
)

// 1. var (...) 只是批量声明变量的语法糖，不是对象！
// 这里定义的 name 和 version 是两个独立的全局变量
var (
	globalName    = "Go"
	globalVersion = 1.24
)

// 2. 真正的“对象”定义：使用 struct（结构体）
// 类似于 Java 的 Class 或 TS 的 Interface
type Language struct {
	Name    string  `json:"name"`    // 导出字段（大写开头），支持 JSON 序列化
	Version float64 `json:"version"` // `json:"..."` 是 Tag，用于定义 JSON 字段名
	Website string  `json:"website"`
}

func main() {
	fmt.Println("=== 1. var (...) 变量组 ===")
	// 获取属性：直接使用变量名
	fmt.Printf("Name: %s, Version: %.2f\n", globalName, globalVersion)

	fmt.Println("\n=== 2. Struct 结构体 (Go 中的对象) ===")
	// 定义并初始化对象（使用花括号 {}）
	// 注意：Key 后面是冒号 :
	goLang := Language{
		Name:    "Go",
		Version: 1.24,
		Website: "https://go.dev",
	}

	// 获取属性：使用点号 .
	fmt.Printf("Object: %+v\n", goLang)
	fmt.Printf("Name: %s\n", goLang.Name)
	fmt.Printf("Version: %.2f\n", goLang.Version)

	// 演示转为 JSON 字符串
	jsonData, _ := json.Marshal(goLang)
	fmt.Printf("JSON String: %s\n", string(jsonData))

	fmt.Println("\n=== 3. Map (类似字典/Hash) ===")
	// 另一种类似 JSON 对象的数据结构：Map
	// 适用于 Key-Value 结构，但 Value 类型必须一致（或使用 interface{}）
	config := map[string]interface{}{
		"env":   "dev",
		"debug": true,
		"port":  8080,
	}

	// 获取属性：使用方括号 []
	fmt.Printf("Map: %+v\n", config)
	fmt.Printf("Env: %s\n", config["env"])
}
