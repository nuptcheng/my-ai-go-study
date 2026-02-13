package main // 1. package: 定义包名

import ( // 2. import: 导入包
	"fmt"
	"time"
)

// 3. const: 定义常量
const AppName = "KeywordsDemo"

// 4. type: 定义类型
// 5. struct: 定义结构体
type User struct {
	ID   int
	Name string
}

// 6. interface: 定义接口
type Greeter interface {
	Greet()
}

func (u User) Greet() {
	fmt.Printf("Hello, I am %s\n", u.Name)
}

// 7. func: 定义函数
// 8. return: 从函数返回
func main() {
	// 9. var: 定义变量
	var count int = 5

	fmt.Printf("Welcome to %s\n", AppName)

	// 10. map: 定义映射，map[K]V 表示键类型为 K，值类型为 V，map[1] = "Alice"
	users := map[int]string{
		1: "Alice",
		2: "Bob",
	}

	// 11. range: 遍历 map/slice/channel
	for id, name := range users { // 12. for: 循环
		fmt.Printf("ID: %d, Name: %s\n", id, name)
	}

	// Flow Control Demo
	flowControlDemo(count)

	// Concurrency Demo
	concurrencyDemo()

	// Defer Demo
	deferDemo()
}

func flowControlDemo(n int) {
	fmt.Println("\n--- Flow Control Demo ---")
	
	// 13. if, 14. else: 条件判断
	if n > 0 {
		fmt.Println("Count is positive")
	} else {
		fmt.Println("Count is non-positive")
	}

	// 15. switch, 16. case, 17. default: 多路分支
	switch n {
	case 1:
		fmt.Println("One")
	case 5:
		fmt.Println("Five")
		// 18. fallthrough: 强制执行下一个 case
		fallthrough
	case 6:
		fmt.Println("Six (reached via fallthrough)")
	default:
		fmt.Println("Other")
	}

	// 19. break, 20. continue: 循环控制
	for i := 0; i < 5; i++ {
		if i == 2 {
			continue // 跳过本次循环
		}
		if i == 4 {
			break // 退出循环
		}
		fmt.Print(i, " ")
	}
	fmt.Println()

	// 21. goto: 跳转 (不推荐频繁使用)
	i := 0
Start:
	if i < 2 {
		fmt.Println("Goto loop:", i)
		i++
		goto Start
	}
}

func concurrencyDemo() {
	fmt.Println("\n--- Concurrency Demo ---")
	
	// 22. chan: 定义通道
	c := make(chan string)

	// 23. go: 启动协程
	go func() {
		time.Sleep(100 * time.Millisecond)
		c <- "Done"
	}()

	// 24. select: 监听通道操作
	select {
	case msg := <-c:
		fmt.Println("Received:", msg)
	case <-time.After(200 * time.Millisecond):
		fmt.Println("Timeout")
	}
}

func deferDemo() {
	fmt.Println("\n--- Defer Demo ---")
	
	// 25. defer: 延迟执行 (LIFO)
	defer fmt.Println("Defer 1: End of function")
	defer fmt.Println("Defer 2: This prints before Defer 1")
	
	fmt.Println("Main logic executing...")
}
