/*
 * @Author: BecomeBamboo
 * @Date: 2026-02-12 16:03:03
 * @LastEditTime: 2026-02-12 16:03:41
 * @Description: 
 */
package main

import "fmt"

// Go 语言的“分号插入规则”决定了左花括号 { 不能单独换行
// 如果你写成:
// func main()
// {
// 编译器会在 func main() 后面自动插入分号，变成 func main();
// 导致 syntax error: unexpected semicolon or newline before {

func main() { // 正确：{ 必须在同一行
	fmt.Println("Hello, World!")

	// 控制结构也是同理
	if true { // 正确
		fmt.Println("True")
	}

	/* 错误写法演示：
	if true
	{
		fmt.Println("Error")
	}
	*/
}
