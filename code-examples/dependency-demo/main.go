package main

import (
	"fmt"
	"strings"

	// 引入外部依赖
	"github.com/google/uuid"
)

func main() {
	// 使用外部依赖生成 UUID
	id := uuid.New()
	
	fmt.Println("Generated UUID:")
	fmt.Printf("  String: %s\n", id.String())
	fmt.Printf("  Upper : %s\n", strings.ToUpper(id.String()))
}
