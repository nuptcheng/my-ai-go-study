
# 2026-02-12 学习记录

[⬅️ 返回主页](../../README.md)

## 📝 2026-02-12 学习总结

### 核心知识点回顾
1. **短变量声明 (`:=`)**: 只能在函数内使用，自动推断类型。
2. **结构体 (`struct`)**: Go 的对象定义方式，通过 `type Name struct {...}` 定义。
3. **语法细节**: 左大括号 `{` 必须位于行尾（由于自动分号插入机制）。
4. **关键字**: 25 个关键字分类记忆（声明、复合、控制、并发、特殊）。
5. **常量**: 无类型常量 (`Untyped Constant`) 具有高度灵活性。
6. **接口 (`interface`)**:
    - **隐式实现 (Duck Typing)**: 不需要显式声明 implements。
    - **多态**: 接口变量存储具体类型，运行时动态派发。
    - **空接口 (`interface{}`)**: 可接收任何类型。
7. **数据类型**: 基本类型 (bool/int/string...) 与 复合类型 (slice/map/chan...)。
8. **切片 (Slice)**: 动态数组，底层由指针、长度、容量组成；扩容机制随版本优化。
9. **工程规范**: 文件名全小写下划线，常量/变量驼峰式；推荐 Standard Go Project Layout。
10. **匿名函数与闭包**:
    - **匿名函数**: 没有名字的函数，常用于 defer、goroutine、callback。
    - **闭包 (Closure)**: 引用了外部变量的函数值，变量生命周期随闭包延长。
11. **数组 (Array)**:
    - **定长**: 长度是类型的一部分 (`[3]int` != `[4]int`)。
    - **值传递**: 函数传参会拷贝整个数组，建议传指针或切片。
    - **初始化**: 支持 `...` 自动推断长度；支持多维数组。

### 今日代码产出
| 演示项目 | 描述 |
| :--- | :--- |
| [var-decl-demo](../../code-examples/var-decl-demo) | 短变量声明与包级变量对比 |
| [struct-demo](../../code-examples/struct-demo) | 结构体定义与对象属性访问 |
| [syntax-demo](../../code-examples/syntax-demo) | 左大括号位置与分号插入规则 |
| [keywords-demo](../../code-examples/keywords-demo) | 25 个关键字实战用法 |
| [const-demo](../../code-examples/const-demo) | 无类型常量与有类型常量对比 |
| [interface-demo](../../code-examples/interface-demo) | 接口隐式实现、多态与类型断言 |
| [types-demo](../../code-examples/types-demo) | Go 全数据类型定义与使用综合演示 |
| [slice-demo](../../code-examples/slice-demo) | 切片底层结构、扩容机制与内存陷阱 |
| [layout-demo](../../code-examples/layout-demo) | 符合 Go 标准规范的项目目录结构演示 |
| [closure-demo](../../code-examples/closure-demo) | 匿名函数与闭包的捕获机制演示 |
| [array-demo](../../code-examples/array-demo) | 数组对象存储、定长特性、多维数组与值传递演示 |

