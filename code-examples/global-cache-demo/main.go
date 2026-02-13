/*
 * @Author: BecomeBamboo
 * @Date: 2026-02-13 11:02:48
 * @LastEditTime: 2026-02-13 11:11:08
 * @Description: 
 */
package main

import (
	"fmt"
	"global-cache-demo/pkg/formulas" // 引入我们定义的包
	"sync"
	"time"
)

func main() {
	// 1. 程序启动阶段：初始化配置
	// 必须在任何业务逻辑开始前完成
	formulas.InitFromRedis()

	fmt.Println("=== 业务逻辑开始 ===")

	// 2. 模拟多处并发使用
	var wg sync.WaitGroup
	wg.Add(3)

	// 模拟场景 A：定时任务模块读取 Cron
	go func() {
		defer wg.Done()
		fmt.Println(" -> [CronJob] 正在读取配置...")
		
		// 直接访问全局变量（类似 JS 的 arr[0]）
		for i, item := range formulas.GlobalCache {
			fmt.Printf("    [CronJob] 任务%d: 时间=%s, 测点=%s\n", i, item.Cron, item.ResultStr)
		}
	}()

	// 模拟场景 B：计算引擎模块读取公式
	go func() {
		defer wg.Done()
		time.Sleep(100 * time.Millisecond) // 模拟稍后执行
		fmt.Println(" -> [CalcEngine] 正在获取计算公式...")
		
		// 通过索引访问（带越界检查的 Getter）
		if dto := formulas.GetByIndex(0); dto != nil {
			fmt.Printf("    [CalcEngine] 正在计算: %s\n", dto.FormulaStr)
		}
	}()

	// 模拟场景 C：API 查询接口
	go func() {
		defer wg.Done()
		time.Sleep(200 * time.Millisecond)
		fmt.Println(" -> [API] 用户查询所有配置")
		// 只读操作在并发下是绝对安全的，不需要加锁
		fmt.Printf("    [API] 返回了 %d 条数据给前端\n", len(formulas.GlobalCache))
	}()

	wg.Wait()
	fmt.Println("=== 所有任务完成 ===")
}
