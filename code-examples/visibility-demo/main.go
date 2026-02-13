/*
 * @Author: BecomeBamboo
 * @Date: 2026-02-13 10:51:21
 * @LastEditTime: 2026-02-13 10:59:00
 * @Description: 
 */
package main

import (
	"fmt"
	"visibility-demo/mypkg"
	// ❌ 为什么不能像 JS 那样 import { NewUser } from ... ?
	// Go 语言的设计哲学是 "显式优于隐式"。
	// 1. mypkg.NewUser 让读者一眼就能看出 NewUser 是来自哪个包，而不是当前文件定义的。
	// 2. 如果非要省略包名，可以使用 "点导入" (Dot Import)，但极其不推荐（容易命名冲突）：
	//    import . "visibility-demo/mypkg" -> 此时可以直接调用 NewUser()
	// 3. 如果包名太长，推荐使用 "别名导入" (Alias Import):
	//    import m "visibility-demo/mypkg" -> m.NewUser()
)

func main() {
	fmt.Println("=== Go 可见性 (Visibility) 演示 ===")

	// 1. 创建对象
	// 语法拆解:
	// mypkg  -> 包名 (package name)
	// .      -> 访问符
	// User   -> 类型名 (struct Type)，注意这里 User 不是方法，而是我们在 mypkg 里定义的 type User struct
	// &      -> 取地址，返回一个指向该对象的指针 (*mypkg.User)
	// {}     -> 初始化字面量
	u := &mypkg.User{
		Username: "Alice",
	}
	fmt.Println("User:", u.Username)

	// 2. 访问字段
	// fmt.Println(u.age) // 编译错误：u.age undefined (cannot refer to unexported field or method age)
	
	// 3. 正确做法：通过工厂函数创建，通过 Getter 方法访问
	u2 := mypkg.NewUser("Bob", 25)
	fmt.Printf("User2: %s, Age: %d (通过 GetAge() 访问)\n", u2.Username, u2.GetAge())

	// 4. 函数的可见性
	// mypkg.checkAge(10) // 编译错误：cannot refer to unexported name mypkg.checkAge
}
