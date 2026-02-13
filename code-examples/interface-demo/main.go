/*
 * @Author: BecomeBamboo
 * @Date: 2026-02-12 16:19:28
 * @LastEditTime: 2026-02-12 17:00:28
 * @Description: 
 */
package main

import "fmt"

// 1. 定义接口
// 接口只定义“做什么” (Method Signature)，不定义“怎么做” (Implementation)
type Animal interface {
	Speak() string
	Move()
}

// 2. 实现接口 - Dog
type Dog struct {
	Name string
}

// Dog 实现了 Speak 方法
func (d Dog) Speak() string {
	return "Woof!"
}

// Dog 实现了 Move 方法
func (d Dog) Move() {
	fmt.Printf("%s runs on 4 legs\n", d.Name)
}

// 3. 实现接口 - Bird
type Bird struct {
	Name string
}

func (b Bird) Speak() string {
	return "Chirp!"
}

func (b Bird) Move() {
	fmt.Printf("%s flies in the sky\n", b.Name)
}

// 4. 多态 (Polymorphism)
// 这个函数接收任何实现了 Animal 接口的类型
func LetAnimalDoThings(a Animal) {
	fmt.Println("Animal says:", a.Speak())
	a.Move()
}

// 5. 空接口 (Empty Interface) - 类似 Java 的 Object
// interface{} 没有定义任何方法，所以任何类型都实现了它
func PrintAnything(v interface{}) {
	fmt.Printf("Value: %v, Type: %T\n", v, v)
}

func main() {
	// 实例化具体类型
	dog := Dog{Name: "Buddy"}
	bird := Bird{Name: "Tweety"}

	fmt.Println("--- 多态演示 ---")
	// Go 的接口是“隐式实现”的
	// 我们没有声明 Dog implements Animal，只要方法签名匹配，就是实现了
	LetAnimalDoThings(dog)
	LetAnimalDoThings(bird)

	fmt.Println("\n--- 接口切片 ---")
	// 可以把不同类型的对象放在同一个接口切片中
	zoo := []Animal{dog, bird}
	for _, animal := range zoo {
		fmt.Println(animal.Speak())
	}

	fmt.Println("\n--- 空接口与类型断言 ---")
	PrintAnything(42)
	PrintAnything("Hello")
	PrintAnything(dog)

	// 类型断言 (Type Assertion)
	var anyVal interface{} = "Hello Interface"
	str, ok := anyVal.(string) // 尝试转回 string
	if ok {
		fmt.Println("It's a string:", str)
	} else {
		fmt.Println("It's not a string")
	}
}
