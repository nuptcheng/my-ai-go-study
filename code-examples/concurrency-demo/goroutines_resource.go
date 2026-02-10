package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	// 强制 GC 以获取干净的初始状态
	runtime.GC()
	var m1, m2 runtime.MemStats
	runtime.ReadMemStats(&m1)

	numGoroutines := 100000
	fmt.Printf("开始创建 %d 个协程...\n", numGoroutines)

	var wg sync.WaitGroup
	wg.Add(numGoroutines)
	
	// 用于阻塞协程的通道
	done := make(chan struct{})

	start := time.Now()
	for i := 0; i < numGoroutines; i++ {
		go func() {
			wg.Done()     // 通知主线程“我已启动”
			<-done        // 阻塞，保持协程存活
		}()
	}

	// 等待所有协程启动
	wg.Wait()
	duration := time.Since(start)

	// 再次读取内存
	runtime.ReadMemStats(&m2)

	// 计算差值
	totalAlloc := m2.TotalAlloc - m1.TotalAlloc
	sysMem := m2.Sys - m1.Sys // 从 OS 申请的内存变化
	heapInUse := m2.HeapInuse - m1.HeapInuse
	
	// 粗略计算每个协程的内存消耗
	// 注意：HeapInuse 包含了协程栈（在 Go 中栈分配在堆上或由运行时管理），
	// 但 Sys 更能反映进程视角的内存增长。
	// 这里主要看 Sys 或 HeapInuse 的增量。
	
	fmt.Printf("创建耗时: %v\n", duration)
	fmt.Printf("内存变化 (Sys): %d MB\n", bToMb(sysMem))
	fmt.Printf("内存变化 (HeapInuse): %d MB\n", bToMb(heapInUse))
	fmt.Printf("平均每个协程占用: ~%.2f KB\n", float64(sysMem)/float64(numGoroutines)/1024)

	// 释放协程
	close(done)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
