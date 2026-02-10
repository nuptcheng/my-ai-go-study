package main

import (
	"fmt"
	"project-layout-demo/internal/logic"
	"project-layout-demo/pkg/utils"
)

func main() {
	fmt.Println("App Starting...")
	
	// 使用 pkg 下的通用工具
	utils.Help()
	
	// 使用 internal 下的业务逻辑
	res := logic.DoBusiness()
	fmt.Println("Result:", res)
}
