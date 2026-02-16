package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("=== 1. 定义对象切片 ===")
	// 方式 A: 声明 nil 切片 (类似 null)
	var listA []Person
	fmt.Printf("listA: %v, len: %d, cap: %d, isNil: %v\n", listA, len(listA), cap(listA), listA == nil)

	fmt.Println("\n=== 2. 使用 make 创建切片 ===")
	// make([]Type, len, cap)
	// len=2: 初始有两个默认零值的元素
	// cap=5: 底层数组容量为5，后续append 3个元素内不会触发扩容
	listB := make([]Person, 2, 5)
	listB[0] = Person{Name: "Alice", Age: 18}
	listB[1] = Person{Name: "Bob", Age: 20}
	fmt.Printf("listB: %v, len: %d, cap: %d\n", listB, len(listB), cap(listB))

	fmt.Println("\n=== 3. 切片初始化 (字面量) ===")
	listC := []Person{
		{Name: "Charlie", Age: 22}, // 注意: Go 允许省略类型名 Person
		{Name: "Dave", Age: 25},
		{Name: "Eve", Age: 30},
	}
	fmt.Printf("listC: %v\n", listC)

	fmt.Println("\n=== 4. 切片截取 (Slicing) [start:end] ===")
	// 规则: 左闭右开 [start, end)
	// listC 有 3 个元素: 0:Charlie, 1:Dave, 2:Eve
	subList := listC[1:3] // 取 index 1 和 2 (不包含 3)
	fmt.Printf("原始 listC: %v\n", listC)
	fmt.Printf("截取 listC[1:3]: %v (Dave, Eve)\n", subList)
	
	// 省略写法
	fmt.Printf("listC[:2]: %v (从头到2)\n", listC[:2])
	fmt.Printf("listC[1:]: %v (从1到尾)\n", listC[1:])

	fmt.Println("\n=== 5. len() 和 cap() 的区别 ===")
	// subList 是基于 listC 的底层数组
	// len: 当前切片里的元素个数 (Dave, Eve -> 2个)
	// cap: 从切片的 start 位置到底层数组末尾的个数 (Dave(1) 到 Eve(2) -> 2个)
	fmt.Printf("subList len: %d, cap: %d\n", len(subList), cap(subList))

	fmt.Println("\n=== 6. append() 方法 ===")
	// 向 listB (len=2, cap=5) 追加元素
	// 因为 cap 够用，不会换底层数组
	listB = append(listB, Person{Name: "Frank", Age: 35})
	fmt.Printf("append后 listB: len=%d, cap=%d, %v\n", len(listB), cap(listB), listB)

	// 继续追加，超过 cap 后会触发扩容
	listB = append(listB, Person{Name: "Grace", Age: 40}, Person{Name: "Heidi", Age: 45}, Person{Name: "Ivan", Age: 50})
	fmt.Printf("扩容后 listB: len=%d, cap=%d (容量翻倍/扩容)\n", len(listB), cap(listB))

	fmt.Println("\n=== 7. copy() 方法 ===")
	// 必须先初始化目标切片的长度 (len)，否则 copy 不进去
	dest := make([]Person, len(listC))
	count := copy(dest, listC) // 返回复制的元素个数
	fmt.Printf("复制了 %d 个元素\n", count)
	fmt.Printf("源 listC: %v\n", listC)
	fmt.Printf("新 dest : %v (修改它不会影响 listC)\n", dest)
	
	// 修改 dest 验证独立性
	dest[0].Name = "CopyCat"
	fmt.Printf("修改dest后: %v\n", dest[0].Name)
	fmt.Printf("原始listC: %v (未变)\n", listC[0].Name)
}
