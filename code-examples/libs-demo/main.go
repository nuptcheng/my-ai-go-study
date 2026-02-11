package main

import (
	"fmt"
	"log"

	"github.com/antonmedv/expr"
)

// 这个 Demo 演示了如何使用 Go 处理动态公式和逻辑运算
// 场景：电商系统根据用户等级和订单金额动态计算折扣
// 这种逻辑通常写在数据库配置里，而不是硬编码在代码里

func main() {
	// 1. 定义上下文环境（输入参数）
	// 这是一个 map，key 是变量名，value 是变量值
	env := map[string]interface{}{
		"order_amount": 100, // 订单金额
		"vip_level":    2,   // VIP 等级
		"user_tag":     "new_user",
	}

	// 2. 定义动态公式（支持三元运算）
	// 逻辑：
	// - 如果是 VIP (等级 > 1)，打 8 折
	// - 如果是新用户 (tag == 'new_user')，减 10 块
	// - 否则原价
	// 注意：Go 语言本身不支持 '?' 三元运算符，但 expr 库完美支持
	script := `vip_level > 1 ? order_amount * 0.8 : (user_tag == "new_user" ? order_amount - 10 : order_amount)`

	fmt.Printf("环境参数: %+v\n", env)
	fmt.Printf("计算公式: %s\n", script)

	// 3. 编译公式 (Compile)
	// 这一步会检查语法错误，并生成高效的字节码
	program, err := expr.Compile(script, expr.Env(env))
	if err != nil {
		log.Fatalf("公式编译失败: %v", err)
	}

	// 4. 运行公式 (Run)
	output, err := expr.Run(program, env)
	if err != nil {
		log.Fatalf("公式执行失败: %v", err)
	}

	// 5. 输出结果
	fmt.Printf("计算结果: %v (类型: %T)\n", output, output)

	// 演示另一个例子：简单的布尔逻辑判断
	// 场景：判断是否允许发货 (库存 > 0 且 账户余额 >= 订单金额)
	stockEnv := map[string]interface{}{
		"stock":   5,
		"balance": 50,
		"price":   100,
	}
	rule := `stock > 0 && balance >= price`
	
	allowed, _ := expr.Eval(rule, stockEnv)
	fmt.Printf("\n发货规则: %s, 环境: %+v -> 允许发货? %v\n", rule, stockEnv, allowed)
}
