/*
 * @Author: BecomeBamboo
 * @Date: 2026-02-12 15:58:12
 * @LastEditTime: 2026-02-12 16:03:24
 * @Description: 
 */
package main

import "fmt"

// 1. 函数外（包级变量）：必须使用 var 关键字
// 错误写法: globalVal := 10 (编译报错: non-declaration statement outside function body)

var globalInt = 100            // 完整声明（自动推导类型）
var globalString string = "Hi" // 显式类型声明

// 也可以用 var 块批量声明
var (
	name    = "Go"
	version = 1.24
)

func main() {
	// 2. 函数内：推荐使用短变量声明 :=
	// 含义：声明 + 初始化 + 类型推导
	// 相当于: var message string = "Hello"
	message := "Hello, World!"

	// 这里的 version 是局部变量，遮蔽了上面的包级变量 version
	version := 2.0

	fmt.Printf("Global Int: %d\n", globalInt)
	fmt.Printf("Message: %s\n", message)
	fmt.Printf("Version: %.1f (Local)\n", version)

	// 注意：:= 左边至少要有一个新变量
	// message := "New" // 编译错误: no new variables on left side of :=
	message, status := "New", "OK" // 合法，status 是新变量
	fmt.Println(message, status)
}
