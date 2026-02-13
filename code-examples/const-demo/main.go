package main

import "fmt"

// 1. 无类型常量 (Untyped Constant)
// 这里的 AppName 没有被显式定义为 string 类型
// 它是一个 "无类型字符串常量"，可以被隐式转换为任何兼容的字符串类型
const AppName = "KeywordsDemo"

// 2. 有类型常量 (Typed Constant)
// 显式指定了类型为 string
const Version string = "1.0.0"

// 定义一个自定义类型
type MyString string

func main() {
	fmt.Printf("AppName type: %T, value: %v\n", AppName, AppName)

	// --- 演示区别 ---

	// 场景 A: 赋值给标准 string
	var s1 string = AppName // ✅ OK: 无类型常量自动转为 string
	var s2 string = Version // ✅ OK: 类型匹配
	fmt.Println(s1, s2)

	// 场景 B: 赋值给自定义类型 MyString
	var m1 MyString = AppName // ✅ OK: 无类型常量 "KeywordsDemo" 自动转为 MyString
	
	// var m2 MyString = Version 
	// ❌ 编译错误: cannot use Version (untyped string constant "1.0.0") as MyString value in variable declaration
	// 修正: Version 已经被锁死为 string 类型，不能直接赋值给 MyString，需要强转
	var m2 MyString = MyString(Version)

	fmt.Printf("m1: %v (type: %T)\n", m1, m1)
	fmt.Printf("m2: %v (type: %T)\n", m2, m2)

	// --- 总结 ---
	// 推荐写法: const Name = "Value" (不写类型)
	// 理由: 无类型常量更灵活，可以自动适配不同的目标类型，减少强制类型转换。
}
