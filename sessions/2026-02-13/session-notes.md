# 2026-02-13 学习记录

[⬅️ 返回主页](../../README.md)

## 📝 2026-02-13 学习总结

### 核心知识点回顾
1. **数组 (Array)**: 定长、值类型、无法扩容。适用于小对象或固定长度场景。
2. **指针 (Pointer)**:
    - **语法糖**: 结构体指针 `p.Name` 自动解引用为 `(*p).Name`。
    - **性能**: 小对象传值快，大对象传指针快。
    - **命名**: 推荐名词命名 (`user`) 或 `p`，避免匈牙利命名 (`ptr_user`)。
3. **可见性 (Visibility)**: 首字母大写 Public，小写 Private。规则适用于字段、函数、类型等所有标识符。
4. **包管理 (Packages)**:
    - **包名 vs 文件名**: 包名决定引用方式 (`pkg.Func`)，文件名仅作容器。推荐包名=目录名。
    - **同包冲突**: 同包下跨文件也不能有同名函数/变量。
    - **导入**: 显式导入 `import "fmt"`，避免 JS 式的隐式导入。
5. **工程实践**:
    - **Global Cache**: 包级变量 + `Init()` 实现只读全局缓存，并发安全。
    - **Struct Tag**: `json:"name"` 用于 JSON 映射。
    - **JSON 解析**: 标准库 `encoding/json` 开箱即用。
    - **命名规范**: 允许 `DTO` 等后缀以增强语义。
6. **切片 (Slice)**:
    - **引用类型**: 底层指向数组，包含 `ptr`, `len`, `cap`。
    - **扩容**: `append` 超过 `cap` 时自动扩容（成倍增长）。
    - **截取**: `[start:end]` 左闭右开，共享底层数组。
    - **初始化**: 字面量、`make`、`var` (nil)。

### 今日代码产出
| 演示项目 | 描述 |
| :--- | :--- |
| [array-demo](../../code-examples/array-demo/main.go) | 数组对象存储、定长特性、多维数组与值传递演示 |
| [pointer-demo](../../code-examples/pointer-demo/main.go) | 演示了 `&`/`*` 运算、结构体指针自动解引用、值传递 vs 指针传递性能对比、指针数组与二级指针 |
| [visibility-demo](../../code-examples/visibility-demo/main.go) | 演示了跨包访问权限控制（Public vs Private）及导入语法 |
| [global-cache-demo](../../code-examples/global-cache-demo/pkg/formulas/store.go) | 演示了全局缓存模式、Struct Tag 的使用以及真实 JSON 解析 (`encoding/json`) |
| [slice-ops-demo](../../code-examples/slice-ops-demo/main.go) | 演示了切片定义、make、初始化、截取、len/cap、append/copy 操作 |

---

## 第十一部分：Go 数组 (Array) 全解析

### 1. 数组的核心特性
*   **定长**: 数组的长度是类型的一部分。`[3]int` 和 `[4]int` 是完全不同的类型，不能互相赋值。
*   **值类型**: 数组赋值或传参时，会**拷贝整个数组**。
*   **无法扩容**: 初始化后长度固定，不能 `append`。

### 2. 常见疑问解答
*   **Q: 数组能存对象 (Struct) 吗？**
    *   **A**: 可以。例如 `var people [3]Person`。Go 的数组是内存连续的，存 struct 也是连续存放，性能很高。
*   **Q: 数组大小必须固定吗？**
    *   **A**: 是的。如果需要动态大小，请使用**切片 (Slice)**。JS 的 `[]` 对应 Go 的 Slice。
*   **Q: 如何判断用数组还是切片？**
    *   **切片**: 99% 的业务场景（如 API 返回列表、数据库查询结果）。
    *   **数组**: 长度完全确定且很小（如 MD5 的 16字节、UUID 的 16字节、经纬度坐标），或者对内存布局有极致优化需求。
*   **Q: `...` 的用法？**
    *   **A**: `arr := [...]int{1, 2, 3}` 编译器会自动数数，等价于 `[3]int`。
*   **Q: 多维数组？**
    *   **A**: 支持。如 `[2][3]int` 表示 2 行 3 列的矩阵。
*   **Q: 数组传参怎么用？**
    *   **A**: 直接传 `func(arr [3]int)` 是**值拷贝**（修改不影响原数组）。通常建议传切片 `func(arr []int)` 或指针 `func(arr *[3]int)` 以避免拷贝。

### 代码示例
- [array-demo](../../code-examples/array-demo/main.go): 演示了数组存对象、定长限制、`...` 初始化、二维数组以及值传递特性。

---

## 第十二部分：Go 指针 (Pointer) 详解

### 1. 指针核心概念
*   **指针是什么**: 存储另一个变量内存地址的变量。
*   **符号**:
    *   `&` (取地址): 获取变量在内存中的位置。
    *   `*` (解引用/Dereference): 获取该地址指向的值。

### 2. 常见疑问解答
*   **Q: 结构体指针怎么用？类似 JS 的 `arr[0].name` 吗？**
    *   **A**: 是的，Go 有语法糖。
    *   定义: `p := &Person{Name: "Go"}`
    *   使用: `p.Name` (编译器自动转换为 `(*p).Name`)。这让 JS 开发者感到非常亲切，不需要手动解引用。
    *   数组指针: `people := [2]Person{{...}}`, `ptr := &people[0]`, `ptr.Name` 也是直接访问。

*   **Q: 直接用指针 vs 直接用对象 (值)，哪个快？**
    *   **小对象 (如 int, bool)**: **传值**更快（直接在栈上复制，无 GC 压力）。
    *   **大对象 (如大结构体)**: **传指针**更快（只复制 8 字节地址，避免大数据拷贝）。
    *   **语义**: 如果函数需要**修改**原对象，必须传指针。

*   **Q: 指针变量命名规范？`ptr_number` 还是 `numberPtr`？**
    *   **A**: Go 社区倾向于简洁。
    *   不推荐: Hungarian notation (`ptr_number`, `p_number`).
    *   推荐: 直接用名词 (如 `user`, `config`)，如果必须区分，用 `userPtr` 或单字母 `p` (在短函数中)。
    *   例如: `func NewUser() *User` 返回的就是指针，但接收者通常叫 `u` 或 `user`。

*   **Q: 指针数组 (Array of Pointers) vs 数组指针 (Pointer to Array)？**
    *   **指针数组 (`[3]*int`)**: 数组里存了 3 个地址。常用于对象很大，不想在数组里挪动数据的场景。
    *   **数组指针 (`*[3]int`)**: 指向整个数组的指针。常用于传递定长数组给函数且避免拷贝。

*   **Q: 指向指针的指针 (`**int`)？**
    *   **A**: 较少使用。通常出现在需要修改“指针本身指向哪里”的场景，或者与 C 语言库交互时。

### 3. 重要补充：字段与函数的可见性 (Visibility)
*   **规则统一**: 不仅是结构体字段，**函数 (Function)**、**方法 (Method)**、**变量 (Variable)**、**类型 (Type)**、**常量 (Constant)** 都遵循首字母大小写规则。
*   **Public (大写)**: 如 `NewUser`, `GetAge`。可以被其他包调用。
*   **Private (小写)**: 如 `checkAge`。只能在当前包内部被调用，外部不可见。

### 代码示例
- [pointer-demo](../../code-examples/pointer-demo/main.go): 演示了 `&`/`*` 运算、结构体指针自动解引用、值传递 vs 指针传递性能对比、指针数组与二级指针。
- [visibility-demo](../../code-examples/visibility-demo/main.go): 演示了跨包访问权限控制（Public vs Private）及导入语法。

---

## 第十三部分：Go vs JS 模块导入机制对比

### 1. 核心差异：显式 vs 隐式
*   **JS/TS (`import { NewUser } from ...`)**: 喜欢把东西直接导入到当前命名空间，用的时候不用带前缀。方便，但容易不知道 `NewUser` 到底是从哪来的（除非回去看 import）。
*   **Go (`import "mypkg"`)**: 强制使用 `mypkg.NewUser`。
    *   **优点**: 阅读代码时，一眼就知道 `NewUser` 是 `mypkg` 包里的，而不是当前文件定义的。
    *   **缺点**: 写起来稍微啰嗦一点。

### 2. Go 的几种导入方式
1.  **默认导入**: `import "fmt"` -> 使用 `fmt.Println()`。
2.  **别名导入 (Alias Import)**: `import f "fmt"` -> 使用 `f.Println()`。
    *   场景: 包名太长、包名冲突（如两个包都叫 `json`）。
3.  **点导入 (Dot Import)**: `import . "fmt"` -> 直接使用 `Println()`。
    *   **警告**: 极不推荐在业务代码中使用！会让代码难以阅读（不知道函数哪里来的），且容易产生命名冲突。通常只在**测试代码**（如 Gomega 库）中使用。
4.  **匿名导入 (Blank Import)**: `import _ "database/sql"`。
    *   作用: 不使用包里的函数，只为了执行包的 `init()` 初始化函数（如注册数据库驱动）。

---

## 第十四部分：Struct Tag 与命名规范

### 1. `json:"..."` 是什么？
这叫做 **Struct Tag (结构体标签)**。
*   **作用**: 给字段添加元数据。最常见的用途是控制 JSON 序列化/反序列化。
*   **示例**:
    ```go
    type User struct {
        Name string `json:"user_name"` // Go里叫 Name，转成 JSON 变成 "user_name"
        Age  int    `json:"age,omitempty"` // omitempty: 如果 Age 是 0，生成的 JSON 里就没有这个字段
    }
    ```
*   **原因**: Go 强制要求公开字段首字母大写 (`Name`)，但 JSON API 通常习惯小写 (`name` 或 `user_name`)。Tag 就是它们之间的**翻译官**。

### 2. Go 的命名哲学 vs Java
*   **Java**: 喜欢后缀，如 `UserDto`, `UserVo`, `UserEntity`, `UserService`, `UserImpl`。
*   **Go**: 崇尚**简洁**，通常不推荐后缀。
    *   **但是**: 只要团队内部统一，加上后缀（如 `FormulaDto`）完全没问题，尤其是为了明确区分数据对象和领域对象时。**规则是为人服务的**。

### 3. JSON 解析包
*   Go 标准库内置了强大的 JSON 处理包：`encoding/json`。
*   不需要像 Java 那样引入 Jackson 或 Gson。
*   核心函数：
    *   `json.Unmarshal(data, &obj)`: JSON 字符串 -> Go 对象 (反序列化)
    *   `json.Marshal(obj)`: Go 对象 -> JSON 字符串 (序列化)

### 代码示例
- [global-cache-demo](../../code-examples/global-cache-demo/pkg/formulas/store.go): 演示了全局缓存模式、Struct Tag 的使用以及真实 JSON 解析 (`encoding/json`)。

---

## 第十五部分：包名 (Package) 与文件名 (Filename)

### 1. 核心概念
*   **文件名 (`store.go`)**: 只是一个容器。文件名可以随便取，不影响代码调用。
    *   例如，我可以把 `store.go` 改名为 `data.go`，只要里面的代码没变，外部调用的代码一行都不用改。
*   **包名 (`package formulas`)**: 这才是真正的**命名空间**。
    *   外部调用时，使用的是包名：`formulas.InitFromRedis()`。

### 2. Go 的最佳实践
*   **规则**: 同一个目录下的所有 `.go` 文件，必须声明为**同一个包名**。
*   **建议**: 包名最好和**目录名**保持一致。
    *   目录: `pkg/formulas/`
    *   包名: `package formulas`
    *   这样可以避免混淆（如果目录叫 `a`，包名叫 `b`，引用时用 `import "a"`, 调用时用 `b.Func()`，会让人疯掉）。

### 3. 多文件协作
你可以在 `pkg/formulas/` 目录下创建多个文件：
*   `store.go`: 负责存储和数据定义
*   `calc.go`: 负责计算逻辑
*   `utils.go`: 负责辅助函数
只要它们开头都是 `package formulas`，它们就属于同一个包，可以直接互相访问私有变量，无需 import。

### 4. 命名冲突规则
*   **同包唯一性**: 在同一个包 (`package formulas`) 下，**不管分多少个文件**，**绝对不能**有同名的函数、变量、类型或常量。
*   **原因**: 编译器会把这些文件“合并”成一个大文件来看待。
*   **示例**: 如果 `store.go` 定义了 `func Init()`，那么 `calc.go` 里就不能再定义 `func Init()` 了。

---

## 第十六部分：Go 切片 (Slice) 进阶操作

### 1. 定义对象切片 (Slice of Structs)
*   **语法**: `[]Person`。
*   **含义**: 一个动态数组，里面存的全是 `Person` 结构体（值）。
*   **对比 JS**: 就像 JS 的 `[{name: "A"}, {name: "B"}]`。
*   **示例**:
    ```go
    var users []Person // 定义一个 nil 切片
    ```

### 2. 创建切片 (make vs 字面量)
*   **方式 A: 字面量 (推荐用于初始化已知数据)**
    ```go
    users := []Person{
        {Name: "Alice", Age: 18},
        {Name: "Bob", Age: 20},
    }
    ```
*   **方式 B: make (推荐用于动态填充)**
    ```go
    // make([]Type, len, cap)
    // len=0: 初始没有元素
    // cap=10: 预分配10个位置，避免频繁扩容
    users := make([]Person, 0, 10)
    ```

### 3. 切片初始化
*   **nil 切片**: `var s []int` (未分配内存)
*   **空切片**: `s := []int{}` 或 `make([]int, 0)` (已分配内存，但没元素)
*   **带值初始化**: `s := []int{1, 2, 3}`

### 4. 切片截取 (Slicing)
*   **语法**: `slice[start:end]`
*   **规则**: **左闭右开** `[start, end)`。包含 `start`，不包含 `end`。
*   **示例**:
    ```go
    arr := []int{0, 1, 2, 3, 4}
    sub := arr[1:3] // 取 index 1, 2 -> [1, 2]
    ```
*   **省略写法**:
    *   `arr[:2]` -> `arr[0:2]` (从头开始)
    *   `arr[2:]` -> `arr[2:len]` (直到末尾)
    *   `arr[:]` -> 复制整个切片引用

### 5. len() vs cap()
*   **len() (Length)**: 切片里**实际有多少个元素**。
*   **cap() (Capacity)**: 切片**底层数组从 start 开始还有多少空间**。
*   **关系**: `len <= cap`。当 `append` 导致 `len > cap` 时，Go 会自动分配一个更大的底层数组（扩容），并将数据拷贝过去。

### 6. append() 与 copy()
*   **append()**: 向切片末尾追加元素。
    *   **注意**: 必须接收返回值！`s = append(s, val)`。因为扩容后地址会变。
*   **copy()**: 复制两个切片的内容。
    *   **坑**: 目标切片 `dest` 必须有足够的 `len`，否则复制不进去。
    *   **示例**:
        ```go
        src := []int{1, 2}
        dest := make([]int, 2) // len 必须是 2
        copy(dest, src)
        ```

### 代码示例
- [slice-ops-demo](../../code-examples/slice-ops-demo/main.go): 演示了切片定义、make、初始化、截取、len/cap、append/copy 操作。
