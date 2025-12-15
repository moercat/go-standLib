# Go 标准库学习项目

这是一个用于学习和实践 Go 语言标准库及常见用法的集合项目。

## 项目说明

本项目包含各种 Go 语言编程实践案例，涵盖如下主题：

- 数据结构实现
- 并发编程模式
- 排序算法
- 字符串处理
- 测试用例
- 第三方库使用示例

## 包含的模块

- `atm`: 原子操作示例
- `bloomFilter`: 布隆过滤器实现
- `channel`: Go 通道使用示例
- `do_test`: 基准测试示例
- `draw`: 图形绘制示例
- `fun`: 函数特性演示
- `game`: 游戏相关代码
- `lfmt`: 日志格式化工具
- `list`: 双向链表示例
- `map_hash`: 哈希表相关示例
- `plan`: 规划相关代码
- `plugin`: 插件系统示例
- `quicksort`: 排序算法集合
- `rand_opt`: 随机数性能优化示例
- `ring`: 一致性哈希
- `slice`: 切片操作示例
- `sort_demo`: 排序演示
- `str`: 字符串处理操作
- `sync`: 同步原语示例

## 依赖库

本项目使用以下第三方库：

- github.com/aceld/zinx: 网络服务器框架
- github.com/bits-and-blooms/bitset: 位集操作
- github.com/golang/protobuf: Protocol Buffers 支持
- github.com/mroth/weightedrand: 加权随机选择
- github.com/spf13/cast: 类型转换工具
- gonum.org/v1/plot: 数据可视化绘图库

## 使用方法

```bash
go run main.go
```

或进入具体模块运行：
```bash
cd quicksort
go test -v
```

## 许可证

MIT License