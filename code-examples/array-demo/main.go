/*
 * @Author: BecomeBamboo
 * @Date: 2026-02-13 09:49:01
 * @LastEditTime: 2026-02-13 10:26:16
 * @Description: 
 */
package main

import "fmt"

// 1. 定义一个结构体 (JSON对象)
type Person struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("=== Go 数组 (Array) 全面解析 ===")

	// --- 问题 1: 数组里可以放对象吗？ ---
	fmt.Println("\n--- 1. 数组存放结构体 (对象) ---")
	// 定义一个长度为 3 的 Person 数组
	var people [3]Person
	// 初始化
	people[0] = Person{Name: "Alice", Age: 25}
	people[1] = Person{Name: "Bob", Age: 30}
	// people[2] 默认为零值 {"", 0}

	// 或者直接初始化
	employees := [2]Person{
		{Name: "Charlie", Age: 35},
		{Name: "Dave", Age: 40},
	}
	fmt.Printf("People: %v\n", people)
	fmt.Printf("Employees: %v\n", employees)

	// --- 问题 2 & 3: 大小固定 vs 添加元素 ---
	fmt.Println("\n--- 2 & 3. 数组大小固定，无法动态添加 ---")
	var arr [5]int
	// arr[5] = 100 // ❌ 编译错误：索引越界
	// 数组没有 append 操作，必须在初始化时定好大小，或者操作索引
	arr[0] = 10
	fmt.Println("数组长度:", len(arr), "容量:", cap(arr))
	
	// 对比切片 (JS 的 var arr = [])
	// []int 是类型 (切片)
	// {} 是字面量初始化 (Literal)，表示里面是空的
	// 这就像定义对象 Person{Name: "Go"}，这里是定义切片 []int{}
	// 结果：创建一个长度为 0 的非 nil 切片
	// 初始化一个 空的、非 nil 的切片,推荐var slice []int
	slice := []int{} 
	slice = append(slice, 1) // 切片可以 append
	fmt.Println("切片长度 (动态):", len(slice))

	// --- 问题 4: 数组 vs 切片场景 ---
	fmt.Println("\n--- 4. 场景对比 ---")
	fmt.Println("数组: 长度固定 (如 UUID [16]byte, MD5 [16]byte), 内存连续, 性能微优")
	fmt.Println("切片: 长度动态 (99% 的业务场景), 引用类型")

	// --- 问题 5: 使用 ... 自动推断长度 ---
	fmt.Println("\n--- 5. 使用 [...] 自动推断长度 ---")
	autoLenArr := [...]string{"Go", "Python", "Java"}
	fmt.Printf("autoLenArr 类型: %T, 长度: %d\n", autoLenArr, len(autoLenArr))
	// autoLenArr 是 [3]string 类型，长度固定为 3

	// --- 问题 6: 多维数组 ---
	fmt.Println("\n--- 6. 二维数组 ---")
	// 场景：矩阵、棋盘、像素网格
	var grid [2][3]int 
	grid[0] = [3]int{1, 2, 3}
	grid[1] = [3]int{4, 5, 6}
	fmt.Println("2x3 矩阵:", grid)

	// 遍历二维数组
	for i, row := range grid {
		for j, val := range row {
			fmt.Printf("grid[%d][%d] = %d ", i, j, val)
		}
		fmt.Println()
	}

	// --- 问题 7: 数组作为函数参数 (值传递) ---
	fmt.Println("\n--- 7. 数组传参是值拷贝 (Value Copy) ---")
	nums := [3]int{1, 2, 3}
	modifyArray(nums) // 传递的是副本
	fmt.Println("main函数中的 nums:", nums) // 仍然是 [1, 2, 3]

	modifySlice(nums[:]) // 传递切片 (引用语义)
	fmt.Println("传递切片后的 nums:", nums) // 被修改了 [999, 2, 3]
}

// 数组参数：长度必须匹配，且是值拷贝
func modifyArray(arr [3]int) {
	arr[0] = 999
	fmt.Println("modifyArray 内部:", arr)
}

// 切片参数：引用底层数组
func modifySlice(s []int) {
	s[0] = 999
	fmt.Println("modifySlice 内部:", s)
}
