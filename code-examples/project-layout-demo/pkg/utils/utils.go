package utils

import "fmt"

// Help 是通用工具，pkg 包通常可以被外部 module 引用（如果项目是库）
func Help() {
	fmt.Println("This is a helper function")
}
