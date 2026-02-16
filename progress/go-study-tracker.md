[⬅️ 返回主页](../README.md)

# Go 学习进度追踪

## 概览
- 总问题数：25
- 总学习天数：5
- 最后更新：2026-02-16

## 掌握主题
- 类型分类与数量（按规范分为 3 基本 + 8 复合）
- 基本类型详解（整型、浮点、复数、布尔、字符串、rune/byte）
- 复合类型零值与引用性质（slice/map/chan/func/pointer/interface）
- 切片底层原理（ptr/len/cap）与扩容机制（Go 1.18+ 策略）
- 切片共享内存副作用与 Deep Copy
- Go 命名规范（Snake Case 文件名、Camel Case 变量）
- Standard Go Project Layout（cmd/internal/pkg 目录职能）
- 匿名函数定义与立即执行
- 闭包机制与状态保持（Counter 示例）
- 闭包陷阱：循环变量捕获（Go 1.22 变更）
- 数组 (Array) 特性：定长、值类型、多维支持
- 数组 vs 切片的选择标准
- 指针 (Pointer) 核心：& (取地址) 与 * (解引用)
- 结构体指针的自动解引用（Dot Operator 语法糖）
- 值传递 vs 指针传递的性能与语义差异
- Go 访问权限控制（Visibility）：首字母大小写规则 (Public/Private)
- 包名与文件名的关系 (Package vs File)
- 同包跨文件访问与命名冲突规则 (Name Conflict)
- Struct Tag 用法 (`json:"name"`) 与命名规范 (DTO 后缀)
- 标准库 `encoding/json` 的序列化与反序列化
- 全局配置/缓存模式 (Global Cache Pattern) 实现
- 数组 vs 切片的核心差异（长度为类型一部分；切片为描述符）
- map 的基本用法、存在性判断、键的可比性约束
- 并发模型差异（Go 协程 vs Java 平台线程/虚拟线程）
- Goroutine 轻量级特性（2KB 栈、M:N 调度、低切换成本）
- Go 工程化基础（SDK、Go Modules、交叉编译、标准目录结构）
- Go 测试与调试工具（Delve, go test, gotests）
- 变量声明规则（短变量声明 := 限制与包级 var 声明）
- struct 与 var 块的区别及属性访问
- Go 语法规则（左花括号 { 换行限制与自动分号插入）
- Go 25 个关键字（defer/select/fallthrough 等机制）
- Go 常量类型推断（无类型常量 Untyped Constant 的灵活性）
- Go 接口机制（隐式实现、多态、空接口）

## 知识盲点
- 切片扩容规则、append 的复制行为与性能影响
- 方法集与接口实现、动态类型与断言
- 泛型约束（any、comparable 与自定义约束）、类型参数实践
- nil 的适用类型与行为（接收关闭通道、零值处理）

## 学习计划
- 阅读 Effective Go 并配合标准库示例（io、net/http、time、sync、encoding/json）
- 编写并运行切片与 map 的性能对比与边界行为小实验
- 接口与泛型：实现通用容器与算法（例如泛型集合与排序）
- 并发基础：goroutine、channel、select、context 取消与超时

## 索引
- 问题1-Go 类型一共有几种？说明理由（2026-02-10）
  - 记录位置：../sessions/2026-02-10/session-notes.md#问题1-go-类型一共有几种？说明理由
- 问题2-Go协程与Java线程的区别（优劣势、场景、资源差异对比表格）（2026-02-10）
  - 记录位置：../sessions/2026-02-10/session-notes.md#问题2-go协程与java线程的区别优劣势场景资源差异对比表格
- 问题3-Go开发环境搭建与工程化（SDK版本、依赖管理、交叉编译、调试测试、目录结构）（2026-02-10）
  - 记录位置：../sessions/2026-02-10/session-notes.md#问题3-go开发环境搭建与工程化sdk版本依赖管理交叉编译调试测试目录结构
- 问题4-Go安装后需要在代码里指定路径还是配置全局变量？（2026-02-10）
  - 记录位置：../sessions/2026-02-10/session-notes.md#问题4-go安装后需要在代码里指定路径还是配置全局变量？
- 问题5-那么3种方式建议用哪种安装呢？（brew/pkg/tarball）（2026-02-10）
  - 记录位置：../sessions/2026-02-10/session-notes.md#问题5-那么3种方式建议用哪种安装呢？brewpkgtarball
- 问题6-brew安装后能不能切Go版本，类似nvm list这样？（2026-02-10）
  - 记录位置：../sessions/2026-02-10/session-notes.md#问题6-brew安装后能不能切go版本类似nvm-list这样？
- 问题7-好的，那就用方案B，我该怎么安装然后验证？（2026-02-10）
  - 记录位置：../sessions/2026-02-10/session-notes.md#问题7-好的那就用方案b我该怎么安装然后验证？
- 问题8-Go的依赖包怎么管理，比如Node.js有npm install和node_modules，Go是怎么管理的？（2026-02-10）
  - 记录位置：../sessions/2026-02-10/session-notes.md#问题8-go的依赖包怎么管理比如nodejs有npm-install和node_modulesgo是怎么管理的
- 问题9-常用中间件与计算库推荐 (Redis/InfluxDB/TDengine/MySQL/Expr)（2026-02-10）
  - 记录位置：../sessions/2026-02-10/session-notes.md#第九部分常用中间件与计算库推荐-redisinfluxdbtdenginemysqlexpr
- 问题10-Go语言的短变量声明是什么意思？是只能在函数内部声明是嘛，函数外如何处理？（2026-02-12）
  - 记录位置：../sessions/2026-02-12/session-notes.md#问题10-go语言的短变量声明是什么意思是只能在函数内部声明是嘛函数外如何处理
- 问题11-Go里面定义JSON对象是()而不是{}？，定义后如何获取对象属性呢？（2026-02-12）
  - 记录位置：../sessions/2026-02-12/session-notes.md#第二部分go-中的对象定义与属性访问
- 问题12-Go里面{不能单独一行是吗？（语法规则与分号插入机制）（2026-02-12）
  - 记录位置：../sessions/2026-02-12/session-notes.md#第三部分go-语法规则---左花括号--的位置
- 问题13-Go语言中的25个关键字或保留字请给出说明及代码示例（2026-02-12）
  - 记录位置：../sessions/2026-02-12/session-notes.md#第四部分go-语言-25-个关键字详解
- 问题14-Go常量定义时需要显式指定类型吗？（无类型常量与有类型常量的区别）（2026-02-12）
  - 记录位置：../sessions/2026-02-12/session-notes.md#第五部分go-常量定义与类型推断机制
- 问题15-Go 接口定义、隐式实现与多态使用场景（2026-02-12）
  - 记录位置：../sessions/2026-02-12/session-notes.md#第六部分go-接口定义与使用场景
- 问题16-Go 数据类型全解：基本类型与复合类型分类详解（2026-02-12）
  - 记录位置：../sessions/2026-02-12/session-notes.md#第七部分go-数据类型全解
- 问题17-Go 切片 (Slice) 详解：底层结构、扩容机制与常见陷阱（2026-02-12）
  - 记录位置：../sessions/2026-02-12/session-notes.md#第八部分go-切片-slice-详解与扩容机制
- 问题18-Go 语言编码规范与标准项目结构指南 (Naming & Layout)（2026-02-12）
  - 记录位置：../sessions/2026-02-12/session-notes.md#第九部分go-语言编码规范与标准项目结构
- 问题19-Go 数组 (Array) 全解析：结构体数组、定长特性与多维数组（2026-02-13）
  - 记录位置：../sessions/2026-02-13/session-notes.md#第十一部分go-数组-array-全解析
- 问题20-Go 指针 (Pointer) 详解：符号、结构体访问、传值性能对比与二级指针（2026-02-13）
  - 记录位置：../sessions/2026-02-13/session-notes.md#第十二部分go-指针-pointer-详解
- 问题21-Go 可见性与包管理：首字母规则、包名文件名关系与导入机制（2026-02-13）
  - 记录位置：../sessions/2026-02-13/session-notes.md#第十三部分go-vs-js-模块导入机制对比
- 问题22-Go 全局缓存与 Struct Tag：JSON 解析与命名规范（2026-02-13）
  - 记录位置：../sessions/2026-02-13/session-notes.md#第十四部分struct-tag-与命名规范
- 问题23-Go 切片 (Slice) 进阶操作：定义对象切片、make、截取、len/cap、append/copy（2026-02-13）
  - 记录位置：../sessions/2026-02-13/session-notes.md#第十六部分go-切片-slice-进阶操作
- 问题24-Go 的 for 语句与 range 完整用法：多种形态与遍历细节（2026-02-16）
  - 记录位置：../sessions/2026-02-16/session-notes.md#第一部分go-for-语句与-range-用法
- 问题25-Go map 全解析：初始化、遍历、删除与 Set 模拟（2026-02-16）
  - 记录位置：../sessions/2026-02-16/session-notes.md#第二部分go-map-全解析初始化遍历删除与-set-模拟
