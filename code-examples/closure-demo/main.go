package main

import (
	"fmt"
	"time"
)

// 1. 匿名函数 (Anonymous Function)
// 没有名字的函数，通常在需要临时使用函数的地方定义。

// 2. 闭包 (Closure)
// 闭包是一个函数值，它引用了其函数体之外的变量。
// 该函数可以访问并赋予其引用的变量的值。
// 换句话说，该函数被“绑定”到这些变量。

func main() {
	fmt.Println("=== Go 匿名函数与闭包演示 ===")

	// --- 1. 基础匿名函数 ---
	fmt.Println("\n--- 1. 基础匿名函数 ---")
	
	// 1.1 直接定义并调用 (Immediately Invoked Function Expression - IIFE)
	// func(name string) { ... } 定义了一个函数
	// ("Go") 紧接着表示立即调用这个函数，并将 "Go" 作为参数传递给 name
	func(name string) {
		fmt.Printf("Hello, %s! (直接调用)\n", name)
	}("Go")

	// 1.2 赋值给变量
	greet := func(name string) {
		fmt.Printf("Hello, %s! (变量调用)\n", name)
	}
	greet("Closure")

	// --- 2. 闭包演示 ---
	fmt.Println("\n--- 2. 闭包 (Closure) 演示 ---")
	
	// 创建一个计数器闭包
	// counter 是一个函数，它“捕获”了 seq 变量
	nextNum := intSeq() 

	fmt.Println(nextNum()) // 1
	fmt.Println(nextNum()) // 2
	fmt.Println(nextNum()) // 3

	// 创建另一个独立的计数器，状态互不干扰
	newNextNum := intSeq()
	fmt.Println(newNextNum()) // 1

	// --- 3. 闭包的应用场景 ---
	fmt.Println("\n--- 3. 闭包应用场景: 延迟执行与状态保持 ---")
	
	start := time.Now()
	// 模拟耗时操作
	defer func() {
		fmt.Printf("操作耗时: %v (defer + 闭包捕获 start 变量)\n", time.Since(start))
	}()
	time.Sleep(100 * time.Millisecond)

	// --- 4. 闭包陷阱 (循环变量捕获) ---
	fmt.Println("\n--- 4. 闭包陷阱: 循环变量捕获 ---")
	// 注意：Go 1.22 之前，循环变量 i 是共享的。
	// Go 1.22+ 修复了这个问题，每次迭代都有独立的 i。
	// 为了演示闭包捕获引用的特性，我们这里显式使用一个外部变量。

	var sharedVar int
	funcs := []func(){}
	
	for i := 0; i < 3; i++ {
		sharedVar = i // 模拟旧版本 Go 的行为，或者显式共享变量
		// 闭包捕获的是 sharedVar 的引用，而不是值
		funcs = append(funcs, func() {
			fmt.Printf("捕获的值: %d (地址: %p)\n", sharedVar, &sharedVar)
		})
	}

	fmt.Println("执行闭包组 (大家看到的都是同一个变量的最终值):")
	for _, f := range funcs {
		f()
	}

	// 正确的做法 (如果想要捕获当前值)
	fmt.Println("\n正确做法 (通过参数传递值):")
	funcsCorrect := []func(){}
	for i := 0; i < 3; i++ {
		// 方式1: 重新定义局部变量 (Go 1.22 之前常用)
		// val := i 
		// funcsCorrect = append(funcsCorrect, func() { fmt.Println(val) })

		// 方式2: 作为参数传递 (更通用)
		funcsCorrect = append(funcsCorrect, func(val int) func() {
			return func() {
				fmt.Printf("参数传递的值: %d\n", val)
			}
		}(i))
	}
	for _, f := range funcsCorrect {
		f()
	}
}

// intSeq 返回一个闭包
// 每次调用返回的函数，seq 都会增加
func intSeq() func() int {
	seq := 0
	return func() int {
		seq++
		return seq
	}
}
