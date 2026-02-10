<!--
 * @Author: BecomeBamboo
 * @Date: 2026-01-23 14:09:45
 * @LastEditTime: 2026-02-10 16:21:06
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

- 总问题数：7
- 总学习天数：1
- 最后更新：2026-02-10

---

## 📚 知识点索引

### 语言基础
- [Go 类型一共有几种？说明理由](sessions/2026-02-10/session-notes.md#问题1-go-类型一共有几种？说明理由)

### 并发模型
- [Go协程与Java线程的区别（优劣势、场景、资源差异对比表格）](sessions/2026-02-10/session-notes.md#问题2-go协程与java线程的区别优劣势场景资源差异对比表格)

### 接口与泛型
- 待记录

### 标准库
- 待记录

### 工程化
- [Go开发环境搭建与工程化（SDK版本、依赖管理、交叉编译、调试测试、目录结构）](sessions/2026-02-10/session-notes.md#问题3-go开发环境搭建与工程化sdk版本依赖管理交叉编译调试测试目录结构)
- [Go安装后需要在代码里指定路径还是配置全局变量？](sessions/2026-02-10/session-notes.md#问题4-go安装后需要在代码里指定路径还是配置全局变量？)
- [那么3种方式建议用哪种安装呢？（brew/pkg/tarball）](sessions/2026-02-10/session-notes.md#问题5-那么3种方式建议用哪种安装呢？brewpkgtarball)
- [brew安装后能不能切Go版本，类似nvm list这样？](sessions/2026-02-10/session-notes.md#问题6-brew安装后能不能切go版本类似nvm-list这样？)
- [官方多版本方案 (Scheme B) 安装与验证指南](sessions/2026-02-10/session-notes.md#问题7-好的那就用方案b我该怎么安装然后验证？)

---

## 📅 时间线索引
 
### 2026-02-10
- [Go 类型一共有几种？说明理由](sessions/2026-02-10/session-notes.md#问题1-go-类型一共有几种？说明理由)
- [Go协程与Java线程的区别（优劣势、场景、资源差异对比表格）](sessions/2026-02-10/session-notes.md#问题2-go协程与java线程的区别优劣势场景资源差异对比表格)
- [Go开发环境搭建与工程化（SDK版本、依赖管理、交叉编译、调试测试、目录结构）](sessions/2026-02-10/session-notes.md#问题3-go开发环境搭建与工程化sdk版本依赖管理交叉编译调试测试目录结构)
- [Go安装后需要在代码里指定路径还是配置全局变量？](sessions/2026-02-10/session-notes.md#问题4-go安装后需要在代码里指定路径还是配置全局变量？)
- [那么3种方式建议用哪种安装呢？（brew/pkg/tarball）](sessions/2026-02-10/session-notes.md#问题5-那么3种方式建议用哪种安装呢？brewpkgtarball)
- [brew安装后能不能切Go版本，类似nvm list这样？](sessions/2026-02-10/session-notes.md#问题6-brew安装后能不能切go版本类似nvm-list这样？)
- [官方多版本方案 (Scheme B) 安装与验证指南](sessions/2026-02-10/session-notes.md#问题7-好的那就用方案b我该怎么安装然后验证？)

---

## 快速导航

- [学习进度追踪](progress/go-study-tracker.md) - 查看已掌握的主题和知识盲点
- [所有学习记录](sessions/) - 按日期查看详细的学习记录
- [代码示例](code-examples/) - 学习过程中的代码示例
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
