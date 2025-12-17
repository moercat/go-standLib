# Go StandLib

Go StandLib 是一个 Go 语言标准库学习与实践项目，旨在深入理解和掌握 Go 语言的各种特性和最佳实践。

## 目录

- [项目简介](#项目简介)
- [目录结构](#目录结构)
- [安装](#安装)
- [使用方法](#使用方法)
- [模块说明](#模块说明)
- [测试](#测试)
- [贡献](#贡献)
- [许可证](#许可证)

## 项目简介

本项目是一个专注于 Go 标准库学习的实践集合，通过实现常见的数据结构、算法、并发模式等，帮助开发者更好地理解 Go 语言的核心概念和最佳实践。项目以模块化的方式组织，每个模块都专注特定领域，便于学习和参考。

## 目录结构

```
go-standLib/
├── main.go           # 主程序入口，演示排序功能
├── go.mod            # Go 模块定义
├── go.sum            # Go 模块校验和
├── collections/      # 数据结构模块
│   ├── ap.go         # 链表相关实现
│   ├── ck.go         # 一致性哈希相关
│   ├── consistent_str.go
│   ├── fast_search.go # 快速搜索算法
│   ├── faststr.go    # 字符串优化操作
│   ├── join_str.go   # 字符串连接优化
│   ├── list.go       # 链表数据结构
│   ├── range.go      # 范围类型
│   ├── swiss_map.go  # Swiss Map 实现(高性能 map)
│   ├── two_map.go    # Two Map 实现
│   ├── window.go     # 滑动窗口
│   ├── *.go          # 其他数据结构实现
│   └── *_test.go     # 单元测试
├── concurrent/       # 并发编程模块
│   └── channel_example.go # 通道使用示例
├── utils/            # 工具函数模块
│   ├── quicksort.go  # 快速排序实现
│   ├── mergesort.go  # 归并排序实现
│   ├── heapsort.go   # 堆排序实现
│   ├── string_split.go # 字符串分割优化
│   ├── time.go       # 时间相关工具
│   └── *_test.go     # 单元测试
└── README.md         # 项目说明文档
```

## 安装

```bash
# 克隆项目
git clone https://github.com/username/go-standLib.git
cd go-standLib

# 下载依赖
go mod tidy
```

## 使用方法

运行主程序（演示排序功能）：
```bash
go run main.go
```

单独运行某个模块的测试：
```bash
# 运行 utils 包中的排序算法测试
go test ./utils/quicksort_test.go ./utils/quicksort.go ./utils/mergesort.go

# 运行 collections 包的测试
go test ./collections/...
```

## 模块说明

### collections 模块

该模块包含各种数据结构的实现和优化：

- `list.go` - 链表数据结构实现
- `swiss_map.go` - 高性能的 Swiss Map 实现（基于 Swiss Tables）
- `faststr.go` - 字符串操作优化
- `fast_search.go` - 快速搜索算法
- `window.go` - 滑动窗口数据结构
- `consistent_str.go` - 一致性哈希算法
- `range.go` - 范围类型的实现

### concurrent 模块

该模块展示了 Go 的并发编程实践：

- `channel_example.go` - 通道使用示例

### utils 模块

该模块包含常用的工具函数和算法：

- `quicksort.go` - 快速排序算法实现
- `mergesort.go` - 归并排序算法实现
- `heapsort.go` - 堆排序算法实现
- `string_split.go` - 高效字符串分割算法
- `time.go` - 时间处理相关工具函数

## 测试

运行所有测试：
```bash
go test ./...
```

运行特定模块的测试：
```bash
go test ./utils/...  # 运行 utils 模块的所有测试
go test ./collections/...  # 运行 collections 模块的所有测试
go test ./concurrent/...  # 运行 concurrent 模块的所有测试
```

## 贡献

欢迎提交 Issue 和 Pull Request 来改进此项目。请确保遵循以下指导原则：

1. 在提交代码前，请运行测试确保所有功能正常工作
2. 添加适当的单元测试以验证新功能
3. 保持代码风格一致，遵循 Go 语言的最佳实践
4. 更新相关文档

## 许可证

MIT License