/*
 * @Author: BecomeBamboo
 * @Date: 2026-02-12 18:45:16
 * @LastEditTime: 2026-02-13 08:46:55
 * @Description: 
 */
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("=== Go 切片 (Slice) 详解 ===")

	// 1. 切片基础与底层结构
	fmt.Println("\n--- 1. 切片结构与初始化 ---")
	// 定义一个长度为3，容量为5的切片
	// 切片是对底层数组的一个“视图”
	// 底层结构: struct { ptr *Elem; len int; cap int }
	slice1 := make([]int, 3, 5)
	slice1[0], slice1[1], slice1[2] = 1, 2, 3

	fmt.Printf("slice1: %v (len=%d, cap=%d)\n", slice1, len(slice1), cap(slice1))
	fmt.Printf("slice1 address: %p (底层数组首地址)\n", slice1)

	// 2. 切片扩容演示
	fmt.Println("\n--- 2. 切片扩容机制 (Append) ---")
	// 初始状态
	var s []int
	// 预分配容量可以避免多次扩容带来的性能损耗
	// s = make([]int, 0, 100) 
	
	lastCap := cap(s)
	fmt.Printf("Initial: len=%d, cap=%d, ptr=%p\n", len(s), cap(s), s)

	for i := 0; i < 20; i++ {
		s = append(s, i)
		// 当容量发生变化时，打印日志
		if newCap := cap(s); newCap != lastCap {
			fmt.Printf("扩容发生: len=%2d, cap=%2d -> %2d, ptr=%p (地址改变说明发生了内存重新分配)\n", 
				len(s), lastCap, newCap, s)
			lastCap = newCap
		}
	}
	// 观察扩容规律：
	// Go 1.18 前: <1024 双倍扩容; >=1024 1.25倍
	// Go 1.18 后: 采用了更平滑的扩容策略，不仅看阈值，还会对齐内存分配
	
	// 3. 切片截取与共享底层数组
	fmt.Println("\n--- 3. 切片截取与共享底层数组的坑 ---")
	baseArray := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// 截取: [low : high : max]
	// len = high - low
	// cap = max - low (若未指定 max，则为 原cap - low)
	subSlice := baseArray[2:5] // 包含 index 2,3,4
	
	fmt.Printf("baseArray: %v\n", baseArray)
	fmt.Printf("subSlice : %v (len=%d, cap=%d)\n", subSlice, len(subSlice), cap(subSlice))

	// 修改子切片，底层数组也会变！
	subSlice[0] = 999
	fmt.Println("修改 subSlice[0] = 999 后...")
	fmt.Printf("subSlice : %v\n", subSlice)
	fmt.Printf("baseArray: %v (index 2 被修改了!)\n", baseArray)

	// 4. 切片拷贝 (Deep Copy)
	fmt.Println("\n--- 4. 切片拷贝 (Deep Copy) ---")
	// 如果不想共享底层数组，使用 copy
	copySlice := make([]int, len(subSlice))
	copy(copySlice, subSlice)
	
	copySlice[0] = 888
	fmt.Println("修改 copySlice[0] = 888 后...")
	fmt.Printf("copySlice: %v\n", copySlice)
	fmt.Printf("subSlice : %v (未受影响)\n", subSlice)
	
	// 5. 空切片 vs nil 切片
	fmt.Println("\n--- 5. 空切片 vs nil 切片 ---")
	var nilSlice []int      // 只是声明，未初始化
	emptySlice := []int{}   // 字面量初始化为空
	makeSlice := make([]int, 0) 
	
	fmt.Printf("nilSlice  : len=%d, cap=%d, isNil=%v, ptr=%p\n", len(nilSlice), cap(nilSlice), nilSlice==nil, nilSlice)
	fmt.Printf("emptySlice: len=%d, cap=%d, isNil=%v, ptr=%p (zerobase)\n", len(emptySlice), cap(emptySlice), emptySlice==nil, emptySlice)
	fmt.Printf("makeSlice : len=%d, cap=%d, isNil=%v, ptr=%p (zerobase)\n", len(makeSlice), cap(makeSlice), makeSlice==nil, makeSlice)
	
	// 即使是 nil 切片，append 操作也是安全的
	nilSlice = append(nilSlice, 1)
	fmt.Printf("nilSlice after append: %v\n", nilSlice)
}
