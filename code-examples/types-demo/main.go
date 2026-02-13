package main

import (
	"fmt"
	"math/cmplx"
)

// --- 3. 结构体 (Struct) ---
type Person struct {
	Name string
	Age  int
}

// --- 6. 接口 (Interface) ---
type Greeter interface {
	Greet()
}

func (p Person) Greet() {
	fmt.Printf("Hello, I am %s\n", p.Name)
}

func main() {
	fmt.Println("=== Go 语言数据类型全解 ===")

	// --- 1. 基本类型 (Basic Types) ---
	fmt.Println("\n--- 1. 基本类型 ---")
	
	// 布尔型
	var isGoFun bool = true
	fmt.Printf("Bool: %v (Type: %T)\n", isGoFun, isGoFun)

	// 数值型 - 整数
	var a int = 10
	var b int8 = 127
	var c uint = 10
	fmt.Printf("Int: %d, Int8: %d, Uint: %d\n", a, b, c)

	// 数值型 - 浮点数
	var f1 float32 = 3.14
	var f2 float64 = 3.1415926535
	fmt.Printf("Float32: %f, Float64: %f\n", f1, f2)

	// 数值型 - 复数
	var z complex128 = cmplx.Sqrt(-5 + 12i)
	fmt.Printf("Complex: %v\n", z)

	// 别名类型 (byte = uint8, rune = int32)
	var letter byte = 'A' // ASCII
	var chinese rune = '中' // Unicode
	fmt.Printf("Byte (uint8): %v (%c), Rune (int32): %v (%c)\n", letter, letter, chinese, chinese)

	// 字符串
	var str string = "Hello, Go!"
	fmt.Printf("String: %s\n", str)

	// --- 2. 复合类型 (Composite Types) ---
	fmt.Println("\n--- 2. 复合类型 ---")

	// 1. 数组 (Array) - 固定长度
	var arr [3]int = [3]int{1, 2, 3}
	fmt.Printf("Array: %v (Len: %d)\n", arr, len(arr))

	// 2. 切片 (Slice) - 动态长度
	var slice []int = []int{1, 2, 3}
	slice = append(slice, 4) // 动态追加
	fmt.Printf("Slice: %v (Len: %d, Cap: %d)\n", slice, len(slice), cap(slice))

	// 3. 映射 (Map) - 键值对
	var m map[string]int = map[string]int{"one": 1, "two": 2}
	fmt.Printf("Map: %v\n", m)

	// 4. 结构体 (Struct) - 自定义对象
	var p Person = Person{Name: "Alice", Age: 30}
	fmt.Printf("Struct: %+v\n", p)

	// 5. 指针 (Pointer) - 内存地址
	var ptr *int = &a
	fmt.Printf("Pointer: %v (Value: %d)\n", ptr, *ptr)

	// 6. 函数 (Function) - 一等公民
	add := func(x, y int) int {
		return x + y
	}
	fmt.Printf("Function: Add(2, 3) = %d\n", add(2, 3))

	// 7. 接口 (Interface) - 多态
	var g Greeter = p
	fmt.Print("Interface: ")
	g.Greet()

	// 8. 通道 (Channel) - 并发通信
	ch := make(chan string, 1)
	ch <- "Message"
	msg := <-ch
	fmt.Printf("Channel: Received '%s'\n", msg)
}
