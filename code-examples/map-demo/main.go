package main

import (
	"fmt"
	"time"
)

type RealtimeData struct {
	Value float64
	Ts    time.Time
}

func main() {
	fmt.Println("=== 1. map 初始化方式与容量 ===")
	var nilMap map[string]int
	fmt.Printf("nilMap == nil: %v, len: %d\n", nilMap == nil, len(nilMap))

	literalMap := map[string]int{
		"Alice": 90,
		"Bob":   85,
	}
	fmt.Printf("literalMap: %v, len: %d\n", literalMap, len(literalMap))

	madeMap := make(map[string]int)
	fmt.Printf("madeMap: %v, len: %d\n", madeMap, len(madeMap))

	preAllocated := make(map[string]int, 100)
	fmt.Printf("preAllocated: len: %d (容量通过内部哈希桶管理)\n", len(preAllocated))

	fmt.Println("\n=== 2. 不同键值类型的 map ===")
	stringInt := map[string]int{"k1": 1, "k2": 2}
	fmt.Printf("stringInt: %v\n", stringInt)

	stringString := map[string]string{"env": "prod", "region": "cn"}
	fmt.Printf("stringString: %v\n", stringString)

	nested := map[string]map[string]string{
		"db": {
			"host": "127.0.0.1",
			"port": "3306",
		},
	}
	fmt.Printf("nested: %v\n", nested)

	metrics := map[string]RealtimeData{
		"device-001": {Value: 12.3, Ts: time.Now()},
	}
	fmt.Printf("metrics: %v\n", metrics)

	fmt.Println("\n=== 3. 写入、读取与存在性判断 ===")
	userAges := make(map[string]int)
	userAges["Tom"] = 18
	userAges["Jerry"] = 20
	fmt.Printf("userAges: %v\n", userAges)

	age := userAges["Tom"]
	fmt.Printf("Tom age: %d\n", age)

	ageZero := userAges["Unknown"]
	fmt.Printf("Unknown age (直接取): %d\n", ageZero)

	if v, ok := userAges["Unknown"]; ok {
		fmt.Printf("Unknown exists, age: %d\n", v)
	} else {
		fmt.Println("Unknown 不在 map 中")
	}

	fmt.Println("\n=== 4. 长度与遍历 ===")
	fmt.Printf("len(userAges): %d\n", len(userAges))

	for name, a := range userAges {
		fmt.Printf("range: %s => %d\n", name, a)
	}

	fmt.Println("\n=== 5. 删除元素 ===")
	fmt.Printf("before delete: %v\n", userAges)
	delete(userAges, "Tom")
	fmt.Printf("after delete Tom: %v\n", userAges)
	delete(userAges, "NotExist")
	fmt.Printf("after delete NotExist: %v\n", userAges)

	fmt.Println("\n=== 6. 使用 map 实现 Set 与实时数据缓存 ===")
	set := make(map[string]struct{})
	set["A"] = struct{}{}
	set["B"] = struct{}{}
	fmt.Printf("set: %v\n", set)

	if _, ok := set["A"]; ok {
		fmt.Println("A 在集合中")
	}

	delete(set, "B")
	fmt.Printf("set after delete B: %v\n", set)

	realtime := make(map[string]RealtimeData)
	realtime["device-001"] = RealtimeData{Value: 10.1, Ts: time.Now()}
	realtime["device-001"] = RealtimeData{Value: 11.5, Ts: time.Now()}

	for id, data := range realtime {
		fmt.Printf("device %s => value: %.2f, ts: %s\n", id, data.Value, data.Ts.Format(time.RFC3339))
	}
}

