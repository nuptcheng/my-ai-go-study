/*
 * @Author: BecomeBamboo
 * @Date: 2026-02-13 10:23:49
 * @LastEditTime: 2026-02-13 10:52:55
 * @Description: 
 */
package main

import "fmt"

type Person struct {
	Name     string // 大写开头 = 公有 (Public)，其他包可以访问
	Age      int    // 大写开头 = 公有
	password string // 小写开头 = 私有 (Private)，只有当前包 (main) 能访问
}

func main() {
	fmt.Println("=== Go 指针 (Pointer) 全面解析 ===")

	// --- 问题 3: & 和 * 的基本概念 ---
	fmt.Println("\n--- 3. 运算符基础: & (取地址) 与 * (解引用) ---")
	num := 10
	ptr := &num // & 取出 num 的内存地址
	fmt.Printf("num 的值: %d\n", num)
	fmt.Printf("ptr 的值 (地址): %p\n", ptr)
	fmt.Printf("*ptr 的值 (解引用): %d\n", *ptr) // * 取出地址指向的值

	*ptr = 20 // 修改 ptr 指向的值，也会修改 num
	fmt.Println("修改 *ptr = 20 后, num =", num)

	// --- 问题 1: 结构体指针与属性访问 ---
	fmt.Println("\n--- 1. 结构体指针与类似 JS 的访问体验 ---")
	// 初始化结构体，包括私有字段 password (同一个包内可以访问)
	p1 := Person{Name: "Alice", Age: 30, password: "123"}
	pPtr := &p1

	// Go 的语法糖：虽然 pPtr 是指针，但可以直接用 . 访问属性
	// 编译器会自动把 pPtr.Name 转换为 (*pPtr).Name
	// 两者完全等价！
	fmt.Printf("pPtr.Name:     %s (推荐，Go 语法糖)\n", pPtr.Name)
	fmt.Printf("(*pPtr).Name:  %s (原生写法，完全等价)\n", (*pPtr).Name)
	
	// 关于字段大写：
	// Name, Age 是大写 -> Public (对外可见)
	// password 是小写 -> Private (对外不可见，但本包内可见)
	fmt.Printf("password: %s (本包内可见)\n", pPtr.password)

	// 结构体数组指针 (类似 JS 对象数组)
	people := [2]Person{
		{Name: "Bob", Age: 25},
		{Name: "Charlie", Age: 35},
	}
	// 获取第一个对象的指针
	firstPersonPtr := &people[0]
	fmt.Printf("people[0].Name: %s\n", firstPersonPtr.Name)

	// --- 问题 2: 指针 vs 值拷贝 (性能对比) ---
	fmt.Println("\n--- 2. 性能与场景: 值传递 vs 指针传递 ---")
	bigStruct := Person{Name: "BigData", Age: 100}
	
	// 传值：会拷贝整个结构体 (慢，但安全，不影响原数据)
	passByValue(bigStruct)
	fmt.Println("passByValue 后 Age:", bigStruct.Age) // 没变

	// 传指针：只拷贝地址 (快，8字节，但会修改原数据)
	passByPointer(&bigStruct)
	fmt.Println("passByPointer 后 Age:", bigStruct.Age) // 变了

	// --- 问题 5: 指针数组 ---
	fmt.Println("\n--- 5. 指针数组 (Array of Pointers) ---")
	// 场景：当你需要修改数组里的对象，或者对象很大不想拷贝时
	a, b, c := 1, 2, 3
	ptrArr := [3]*int{&a, &b, &c} // 数组里存的是地址
	*ptrArr[0] = 100 // 修改第一个指针指向的值
	fmt.Println("修改 ptrArr[0] 后, a =", a)

	// --- 问题 6: 指向指针的指针 (二级指针) ---
	fmt.Println("\n--- 6. 指向指针的指针 (**ptr) ---")
	var p *int = &a
	var pp **int = &p // pp 存储了 p 的地址

	fmt.Printf("a = %d\n", a)
	fmt.Printf("p (指向 a) = %p, *p = %d\n", p, *p)
	fmt.Printf("pp (指向 p) = %p, **pp = %d\n", pp, **pp)

	// --- 问题 7: 函数传递指针 ---
	fmt.Println("\n--- 7. 函数传递指针 ---")
	x := 1
	increment(&x) // 显式传递地址
	fmt.Println("increment 后 x =", x)
}

// 值传递：函数内修改无效
func passByValue(p Person) {
	p.Age = 999
}

// 指针传递：函数内修改生效
func passByPointer(p *Person) {
	p.Age = 999
}

func increment(val *int) {
	*val = *val + 1
}
