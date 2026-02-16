# 2026-02-16 学习记录

[⬅️ 返回主页](../../README.md)

## 📝 2026-02-16 学习总结

## 第一部分：Go for 语句与 range 用法

### 1. Go 的 for 有几种形态？

Go 只有一个关键字 for，但可以写出多种常见语言里的循环：

- 经典 for（类似 C/Java）：

```go
for i := 0; i < n; i++ {
    fmt.Println(i)
}
```

- while 循环：

```go
for cond {
    // 当 cond 为假时退出
}
```

- 死循环：

```go
for {
    // 需要用 break 或 return 跳出
}
```

- range 循环（foreach）：

```go
for key, value := range m {
    fmt.Println(key, value)
}
```

### 2. range 的通用模式

range 可以遍历多种类型，返回的含义不同：

- 数组/切片：`for i, v := range slice`，i 是下标，v 是元素的拷贝
- map：`for k, v := range m`，k 是键，v 是值
- 字符串：`for i, r := range s`，按 rune 遍历，i 是字节下标，r 是 Unicode 码点
- channel：`for v := range ch`，从通道不断取值，直到通道被关闭

可以用 `_` 忽略不需要的变量：

```go
for k := range m {
    // 只要 key
}

for _, v := range slice {
    // 只要 value
}
```

### 3. 你的示例：遍历并拷贝 map

```go
for key, value := range oldMap {
    newMap[key] = value
}
```

这段代码的含义是：

- 每次循环，从 oldMap 里取出一个 key、value
- 然后把同样的键值对写入 newMap
- 整体效果就是“浅拷贝一个 map”（value 本身如果是引用类型，还会共享底层数据）

注意：遍历 map 的顺序是不保证的，每次运行的 key 顺序可能不同，所以不能依赖遍历顺序做业务逻辑。

### 4. 和其他语言的区别

- Go 把所有循环统一成 for，没有 while 和 do...while 关键字
- range 是一个通用的“遍历语法”，统一了数组、切片、map、字符串、channel 的遍历方式
- map 的遍历顺序完全不保证，这一点和大多数语言也不同（很多语言的字典在新版本中是有序的）

### 5. 一个常见小坑：v 是拷贝

对切片/数组使用 range 时，v 是元素的拷贝：

```go
nums := []int{1, 2, 3}
for i, v := range nums {
    v = v * 10
    // nums[i] 还是原值，因为 v 是拷贝
}
```

如果要修改原切片，应该用下标回写：

```go
for i := range nums {
    nums[i] = nums[i] * 10
}
```

## 第二部分：Go map 全解析：初始化、遍历、删除与 Set 模拟

### 1. map 初始化与容量

- 使用字面量直接初始化：

```go
scores := map[string]int{
    "Alice": 90,
    "Bob":   85,
}
```

- 使用 make 创建空 map：

```go
ages := make(map[string]int)
```

- 使用 make 预估容量（性能优化）：

```go
cache := make(map[string]int, 1000)
```

- 仅声明不初始化是 nil map，只能读不能写：

```go
var m map[string]int
```

### 2. 不同键值类型的 map

Go 的 map 类型为 `map[KeyType]ValueType`：

- `map[string]int`
- `map[string]string`
- `map[string]map[string]string`
- `map[string]RealtimeData`

键类型必须是可比较的（可以用 `==` 判断），值类型可以是任意类型。

### 3. 写入、读取与存在性判断

```go
userAges := make(map[string]int)
userAges["Tom"] = 18
userAges["Jerry"] = 20

age := userAges["Tom"]

ageZero := userAges["Unknown"]

if v, ok := userAges["Unknown"]; ok {
    fmt.Println("存在", v)
} else {
    fmt.Println("不存在")
}
```

直接读取不存在的 key 会返回值类型的零值，需要通过第二个返回值 ok 区分“不存在”与“值为零”。

### 4. 长度与遍历

```go
fmt.Println(len(userAges))

for name, age := range userAges {
    fmt.Println(name, age)
}
```

map 的遍历顺序不保证，每次运行可能不同。

### 5. 删除元素

```go
delete(userAges, "Tom")
delete(userAges, "NotExist")
```

删除不存在的 key 不会报错。

### 6. 使用 map 实现 Set 与实时数据缓存

用 `map[T]struct{}` 实现集合：

```go
set := make(map[string]struct{})
set["A"] = struct{}{}
set["B"] = struct{}{}

if _, ok := set["A"]; ok {
}

delete(set, "B")
```

用 map 保存实时数据，key 唯一：

```go
type RealtimeData struct {
    Value float64
    Ts    time.Time
}

realtime := make(map[string]RealtimeData)
realtime["device-001"] = RealtimeData{Value: 10.1, Ts: time.Now()}
realtime["device-001"] = RealtimeData{Value: 11.5, Ts: time.Now()}
```

同一个 key 反复写入只是覆盖旧值，非常适合“最新数据缓存”的场景。

### 代码示例

- [map-demo](../../code-examples/map-demo/main.go): 演示 map 初始化、类型组合、写入读取、长度遍历、删除与 Set 模拟。

