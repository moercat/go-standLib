# Go StandLib

> Go 标准库源码学习与本地实验笔记

![Go Version](https://img.shields.io/badge/Go-1.23-blue)
![Status](https://img.shields.io/badge/Status-学习笔记-blue)
![CI](https://img.shields.io/badge/CI-无-lightgrey)

## 定位

**解决了什么问题**：个人在阅读 Go 标准库源码和动手实验过程中的笔记。

**与 [go-algorithms-collection](https://github.com/moercat/go-algorithms-collection) 的区别**：
- `go-standLib` — 源码级实验：SwissTable 实现、range 迭代器、滑动窗口、字符串优化对比
- `go-algorithms-collection` — 刷题笔记：LeetCode 题解 + 常用算法模板

**目标用户**：自己（学习参考用）。

## 内容

```
go-standLib/
├── collections/     ← 数据结构实验
│   ├── swiss_map.go    ← SwissTable 风格哈希表
│   ├── two_map.go      ← 双 map 结构
│   ├── faststr.go      ← 字符串操作优化
│   ├── join_str.go     ← 字符串拼接对比
│   ├── window.go       ← 滑动窗口
│   ├── range.go        ← 范围类型
│   └── list.go         ← 链表
├── concurrent/      ← 并发编程
│   └── channel_example.go
├── utils/           ← 工具函数
│   ├── quicksort.go / mergesort.go / heapsort.go
│   ├── string_split.go
│   └── time.go
└── main.go
```

## 技术栈

| 组件 | 选型 |
|------|------|
| 语言 | Go 1.23 |
| 依赖 | 无外部依赖（仅标准库实验） |

## 状态

| 维度 | 说明 |
|------|------|
| 当前阶段 | 个人学习笔记（持续追加） |
| 最后更新 | 2025-12-17 |
| 活跃度 | 📝 学习笔记（不定期更新） |
