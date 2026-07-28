# Go StandLib

> 🎯 **Go 1.21+ 标准库源码阅读笔记** — 边看边写，边学边实验

> 个人学习 Go 标准库源码的笔记和实验代码。每个笔记对应一个标准库包（含源码摘录 + 自己简化重写 + 思考）。

![Go Version](https://img.shields.io/badge/Go-1.21-blue)
![Notes](https://img.shields.io/badge/笔记-3-blue)
![Status](https://img.shields.io/badge/Status-持续追加-blue)

## ✨ 为什么是这个仓库

- 📖 **笔记 + 实验** — 不是"读完就忘"，是看源码后自己写一遍
- 🎯 **聚焦 1.21+ 新包** — slices / maps / cmp / log/slog ...
- 🔬 **本地实验** — collections/concurrent/utils/ 三个实验目录

## 📚 已完成的笔记

| # | 主题 | 笔记 |
| --- | --- | --- |
| 01 | slices / maps / sort 包 | [notes/01-slices-sort.md](./notes/01-slices-sort.md) |
| 02 | cmp 包 + 内建 min/max | [notes/02-cmp-package.md](./notes/02-cmp-package.md) |
| 03 | log/slog 结构化日志 | [notes/03-log-slog.md](./notes/03-log-slog.md) |

每篇笔记结构：源码摘录 → 自己简化重写 → 我的理解。

## 🗂️ 实验代码

```
go-standLib/
├── notes/           ← 3 篇 1.21+ 新包阅读笔记
├── collections/     ← 数据结构实验（SwissTable / 双 map / 字符串优化）
├── concurrent/      ← 并发编程（channel / goroutine）
├── utils/           ← 工具函数（排序算法 / 字符串 / 时间）
└── main.go
```

## 与 go-algorithms-collection 的区别

- `go-standLib` — **源码级实验**（读标准库 + 自己重写）
- `go-algorithms-collection` — **刷题笔记**（LeetCode 题解 + 算法模板）

## 后续笔记计划

- [ ] testing/slogtest
- [ ] context.AfterFunc
- [ ] unique.Handle