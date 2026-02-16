<!--
 * @Author: BecomeBamboo
 * @Date: 2026-01-23 14:09:45
 * @LastEditTime: 2026-02-16 23:18:56
 * @Description: 
-->
# Go学习知识库

## 项目简介

这是我的Go学习知识库，通过**问题驱动**的方式逐步积累Go知识。

每当我提出一个Go相关的问题，学习助手会：
1. 使用苏格拉底式教学法引导我理解
2. 提供清晰的解释和可运行的代码示例
3. 验证我的理解程度
4. 将问答记录归档到对应日期的session文件
5. 自动更新本README的问题索引

## 学习进度

- 总问题数：16
- 总学习天数：2
- 最后更新：2026-02-12

---

## 📚 知识点索引

### 语言基础
- [Go 类型一共有几种？说明理由](sessions/2026-02-10/session-notes.md#问题1-go-类型一共有几种？说明理由)
- [Go语言的短变量声明是什么意思？是只能在函数内部声明是嘛，函数外如何处理？](sessions/2026-02-12/session-notes.md#第一部分go-短变量声明)
- [Go里面定义JSON对象是()而不是{}？，定义后如何获取对象属性呢？](sessions/2026-02-12/session-notes.md#第二部分go-对象定义-struct)
- [Go里面{不能单独一行是吗？（语法规则与分号插入机制）](sessions/2026-02-12/session-notes.md#第三部分go-语法规则-左大括号)
- [Go语言中的25个关键字或保留字详解与分类记忆](sessions/2026-02-12/session-notes.md#第四部分go-关键字-keywords)
- [Go常量定义时需要显式指定类型吗？（无类型常量与有类型常量的区别）](sessions/2026-02-12/session-notes.md#第五部分go-常量定义与类型推断机制)
- [Go 的 for 语句与 range 用法：三种形态与遍历细节](sessions/2026-02-16/session-notes.md#第一部分go-for-语句与-range-用法)
- [Go map 全解析：初始化、遍历、删除与 Set 模拟](sessions/2026-02-16/session-notes.md#第二部分go-map-全解析初始化遍历删除与-set-模拟)

### 并发模型
- [Go协程与Java线程的区别（优劣势、场景、资源差异对比表格）](sessions/2026-02-10/session-notes.md#问题2-go协程与java线程的区别优劣势场景资源差异对比表格)

### 接口与泛型
- [Go 接口定义、隐式实现与多态使用场景](sessions/2026-02-12/session-notes.md#第六部分go-接口定义与使用场景)
- [Go 数据类型全解：基本类型与复合类型分类详解](sessions/2026-02-12/session-notes.md#第七部分go-数据类型全解)
- [Go 切片 (Slice) 详解：底层结构、扩容机制与常见陷阱](sessions/2026-02-12/session-notes.md#第八部分go-切片-slice-详解与扩容机制)
- [Go 切片 (Slice) 进阶操作：定义对象切片、make、截取、len/cap、append/copy](sessions/2026-02-13/session-notes.md#第十六部分go-切片-slice-进阶操作)
- [Go 语言编码规范与标准项目结构指南 (Naming & Layout)](sessions/2026-02-12/session-notes.md#第九部分go-语言编码规范与标准项目结构)
- [Go 匿名函数与闭包 (Closure)：应用场景与循环变量捕获](sessions/2026-02-12/session-notes.md#第十部分go-匿名函数与闭包-closure)
- [Go 数组 (Array) 全解析](sessions/2026-02-13/session-notes.md#第十一部分go-数组-array-全解析)
- [Go 指针 (Pointer) 详解：符号、结构体访问、传值性能对比与二级指针](sessions/2026-02-13/session-notes.md#第十二部分go-指针-pointer-详解)
  - **重要**: 字段可见性规则 (Public/Private) [示例](code-examples/visibility-demo/main.go)
- [Go 包管理与工程实践：包名/文件名关系、JSON 解析与 Struct Tag](sessions/2026-02-13/session-notes.md#第十五部分包名-package-与文件名-filename)
  - **实战**: 全局缓存/配置加载模式 (Global Cache Pattern) [示例](code-examples/global-cache-demo/main.go)
  - **规范**: 结构体命名 (DTO) 与 JSON Tag [代码](code-examples/global-cache-demo/pkg/formulas/store.go)

### 标准库
- 待记录

### 工程化
- [Go开发环境搭建与工程化（SDK版本、依赖管理、交叉编译、调试测试、目录结构）](sessions/2026-02-10/session-notes.md#问题3-go开发环境搭建与工程化sdk版本依赖管理交叉编译调试测试目录结构)
- [Go安装后需要在代码里指定路径还是配置全局变量？](sessions/2026-02-10/session-notes.md#问题4-go安装后需要在代码里指定路径还是配置全局变量？)
- [那么3种方式建议用哪种安装呢？（brew/pkg/tarball）](sessions/2026-02-10/session-notes.md#问题5-那么3种方式建议用哪种安装呢？brewpkgtarball)
- [brew安装后能不能切Go版本，类似nvm list这样？](sessions/2026-02-10/session-notes.md#问题6-brew安装后能不能切go版本类似nvm-list这样？)
- [官方多版本方案 (Scheme B) 安装与验证指南](sessions/2026-02-10/session-notes.md#问题7-好的那就用方案b我该怎么安装然后验证？)
- [Go依赖管理 (Go Modules) 详解与 Node.js 对比](sessions/2026-02-10/session-notes.md#问题8-go的依赖包怎么管理比如nodejs有npm-install和node_modulesgo是怎么管理的)
- [常用中间件与计算库推荐 (Redis/InfluxDB/TDengine/MySQL/Expr)](sessions/2026-02-10/session-notes.md#第九部分常用中间件与计算库推荐-redisinfluxdbtdenginemysqlexpr)
- [Go 语言编码规范与标准项目结构指南 (Naming & Layout)](sessions/2026-02-12/session-notes.md#第九部分go-语言编码规范与标准项目结构)


---

## 学习记录 (Timeline)

### 2026-02-16
- **文档**: [session-notes.md](sessions/2026-02-16/session-notes.md)

### 2026-02-16 演示代码
- [map-demo](code-examples/map-demo/main.go)

### 2026-02-14
- **文档**: [session-notes.md](sessions/2026-02-14/session-notes.md)

### 2026-02-13
- **今日重点**: 数组、指针、可见性规则、包管理、Struct Tag、全局缓存模式、切片进阶操作。
- **文档**: [session-notes.md](sessions/2026-02-13/session-notes.md)
- **演示代码**:
  - [array-demo](code-examples/array-demo/main.go)
  - [pointer-demo](code-examples/pointer-demo/main.go)
  - [visibility-demo](code-examples/visibility-demo/main.go)
  - [global-cache-demo](code-examples/global-cache-demo/pkg/formulas/store.go)
  - [slice-ops-demo](code-examples/slice-ops-demo/main.go)

### 2026-02-12
- [Go语言的短变量声明是什么意思？是只能在函数内部声明是嘛，函数外如何处理？](sessions/2026-02-12/session-notes.md#第一部分go-短变量声明)
- [Go里面定义JSON对象是()而不是{}？，定义后如何获取对象属性呢？](sessions/2026-02-12/session-notes.md#第二部分go-对象定义-struct)
- [Go里面{不能单独一行是吗？（语法规则与分号插入机制）](sessions/2026-02-12/session-notes.md#第三部分go-语法规则-左大括号)
- [Go语言中的25个关键字或保留字详解与分类记忆](sessions/2026-02-12/session-notes.md#第四部分go-关键字-keywords)
- [Go常量定义时需要显式指定类型吗？（无类型常量与有类型常量的区别）](sessions/2026-02-12/session-notes.md#第五部分go-常量定义与类型推断机制)
- [Go 接口定义、隐式实现与多态使用场景](sessions/2026-02-12/session-notes.md#第六部分go-接口定义与使用场景)
- [Go 数据类型全解：基本类型与复合类型分类详解](sessions/2026-02-12/session-notes.md#第七部分go-数据类型全解)
- [Go 切片 (Slice) 详解：底层结构、扩容机制与常见陷阱](sessions/2026-02-12/session-notes.md#第八部分go-切片-slice-详解与扩容机制)
- [Go 匿名函数与闭包 (Closure)：应用场景与循环变量捕获](sessions/2026-02-12/session-notes.md#第十部分go-匿名函数与闭包-closure)


### 2026-02-10
- [Go 类型一共有几种？说明理由](sessions/2026-02-10/session-notes.md#问题1-go-类型一共有几种？说明理由)
- [Go协程与Java线程的区别（优劣势、场景、资源差异对比表格）](sessions/2026-02-10/session-notes.md#问题2-go协程与java线程的区别优劣势场景资源差异对比表格)
- [Go开发环境搭建与工程化（SDK版本、依赖管理、交叉编译、调试测试、目录结构）](sessions/2026-02-10/session-notes.md#问题3-go开发环境搭建与工程化sdk版本依赖管理交叉编译调试测试目录结构)
- [Go安装后需要在代码里指定路径还是配置全局变量？](sessions/2026-02-10/session-notes.md#问题4-go安装后需要在代码里指定路径还是配置全局变量？)
- [那么3种方式建议用哪种安装呢？（brew/pkg/tarball）](sessions/2026-02-10/session-notes.md#问题5-那么3种方式建议用哪种安装呢？brewpkgtarball)
- [brew安装后能不能切Go版本，类似nvm list这样？](sessions/2026-02-10/session-notes.md#问题6-brew安装后能不能切go版本类似nvm-list这样？)
- [官方多版本方案 (Scheme B) 安装与验证指南](sessions/2026-02-10/session-notes.md#问题7-好的那就用方案b我该怎么安装然后验证？)
- [Go依赖管理 (Go Modules) 详解与 Node.js 对比](sessions/2026-02-10/session-notes.md#问题8-go的依赖包怎么管理比如nodejs有npm-install和node_modulesgo是怎么管理的)
- [常用中间件与计算库推荐 (Redis/InfluxDB/TDengine/MySQL/Expr)](sessions/2026-02-10/session-notes.md#第九部分常用中间件与计算库推荐-redisinfluxdbtdenginemysqlexpr)

---

## 快速导航

- [学习进度追踪](progress/go-study-tracker.md) - 查看已掌握的主题和知识盲点
- [所有学习记录](sessions/) - 按日期查看详细的学习记录
- [代码示例](code-examples/) - 学习过程中的代码示例
- [项目文档](docs/) - 独立的 Go 语言规范与技术文档
- [系统配置](.trae/project_rules.md) - 查看学习系统的配置说明
 - 示例：类型初始化与使用分离 [types-demo](file:///Users/becomebamboo/source/my-ai-study/my-ai-go-study/code-examples/types-demo)
 - 示例：工程目录结构与交叉编译 [project-layout-demo](file:///Users/becomebamboo/source/my-ai-study/my-ai-go-study/code-examples/project-layout-demo)

## 文件结构

```
my-ai-go-study/
├── .trae/                       # 系统配置目录
│   └── project_rules.md         # 学习系统规则
├── README.md                    # 本文件 - 问题索引
├── sessions/                    # 每日学习记录
│   ├── SESSION-TEMPLATE.md      # session记录模板
│   └── YYYY-MM-DD/
│       └── session-notes.md     # 每日问答详细记录
├── progress/                    # 学习进度追踪
│   └── go-study-tracker.md      # 总体进度和学习计划
└── code-examples/               # 代码示例（按需创建）
```

## 使用说明

1. **提问**：直接提出Go相关的问题
2. **学习**：跟随Claude的引导式教学，理解概念
3. **实践**：运行和修改提供的代码示例
4. **复习**：通过本README的问题索引快速找到之前学过的内容
5. **追踪**：查看progress/go-study-tracker.md了解学习进度

## 学习特色

- ✅ **问题驱动**：不预先规划，根据实际需求学习
- ✅ **苏格拉底式教学**：通过提问引导理解，而非直接灌输
- ✅ **代码示例**：每个概念都有可运行的代码示例
- ✅ **自动归档**：所有问题自动索引，方便查找
- ✅ **进度追踪**：清晰了解已掌握和需要加强的领域
- ✅ **权威验证**：技术答案经过在线验证（Effective Go、Go 语言规范、go.dev/blog、pkg.go.dev等）

---

**开始学习**：直接提出你的第一个Go问题吧！
