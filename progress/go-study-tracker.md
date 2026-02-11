# Go 学习进度追踪

## 概览
- 总问题数：8
- 总学习天数：1
- 最后更新：2026-02-10

## 掌握主题
- 类型分类与数量（按规范分为 3 基本 + 8 复合）
- 复合类型零值与引用性质（slice/map/chan/func/pointer/interface）
- 数组 vs 切片的核心差异（长度为类型一部分；切片为描述符）
- map 的基本用法、存在性判断、键的可比性约束
- 并发模型差异（Go 协程 vs Java 平台线程/虚拟线程）
- Goroutine 轻量级特性（2KB 栈、M:N 调度、低切换成本）
- Go 工程化基础（SDK、Go Modules、交叉编译、标准目录结构）
- Go 测试与调试工具（Delve, go test, gotests）

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
  - 记录位置：sessions/2026-02-10/session-notes.md#问题1-go-类型一共有几种？说明理由
- 问题2-Go协程与Java线程的区别（优劣势、场景、资源差异对比表格）（2026-02-10）
  - 记录位置：sessions/2026-02-10/session-notes.md#问题2-go协程与java线程的区别优劣势场景资源差异对比表格
- 问题3-Go开发环境搭建与工程化（SDK版本、依赖管理、交叉编译、调试测试、目录结构）（2026-02-10）
  - 记录位置：sessions/2026-02-10/session-notes.md#问题3-go开发环境搭建与工程化sdk版本依赖管理交叉编译调试测试目录结构
- 问题4-Go安装后需要在代码里指定路径还是配置全局变量？（2026-02-10）
  - 记录位置：sessions/2026-02-10/session-notes.md#问题4-go安装后需要在代码里指定路径还是配置全局变量？
- 问题5-那么3种方式建议用哪种安装呢？（brew/pkg/tarball）（2026-02-10）
  - 记录位置：sessions/2026-02-10/session-notes.md#问题5-那么3种方式建议用哪种安装呢？brewpkgtarball
- 问题6-brew安装后能不能切Go版本，类似nvm list这样？（2026-02-10）
  - 记录位置：sessions/2026-02-10/session-notes.md#问题6-brew安装后能不能切go版本类似nvm-list这样？
- 问题7-好的，那就用方案B，我该怎么安装然后验证？（2026-02-10）
  - 记录位置：sessions/2026-02-10/session-notes.md#问题7-好的那就用方案b我该怎么安装然后验证？
- 问题8-Go的依赖包怎么管理，比如Node.js有npm install和node_modules，Go是怎么管理的？（2026-02-10）
  - 记录位置：sessions/2026-02-10/session-notes.md#问题8-go的依赖包怎么管理比如nodejs有npm-install和node_modulesgo是怎么管理的
