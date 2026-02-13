/*
 * @Author: BecomeBamboo
 * @Date: 2026-02-13 10:51:12
 * @LastEditTime: 2026-02-13 10:56:13
 * @Description: 
 */
package mypkg

// User 结构体首字母大写 -> Public (其他包可见)
type User struct {
	Username string // 大写 -> Public (其他包可访问)
	age      int    // 小写 -> Private (只有 mypkg 包内可访问)
}

// 工厂函数：通常用于创建对象，可以处理私有字段初始化
// NewUser 是大写开头 -> Public (其他包可以调用)
func NewUser(name string, age int) *User {
	// 调用内部私有辅助函数检查数据
	if !checkAge(age) {
		age = 0
	}
	return &User{
		Username: name,
		age:      age, // 包内可以访问 age
	}
}

// 私有函数：小写开头 -> Private (只能在 mypkg 包内部被调用)
// 外部 main 包无法直接调用 mypkg.checkAge()
func checkAge(age int) bool {
	return age >= 0
}

// 公有方法：可以获取私有字段
func (u *User) GetAge() int {
	return u.age
}
