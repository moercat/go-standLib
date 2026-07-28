# Go 1.21+ 标准库新包阅读笔记 #3

> 阅读对象：`log/slog`（Go 1.21 引入的结构化日志）

## 背景

Go 生态里有大量第三方日志库（zap / zerolog / logrus），
1.21 终于引入官方结构化日志 `log/slog`。

## 基础用法

### 默认 logger

```go
import "log/slog"

slog.Info("hello")
slog.Info("user logged in", "user", "alice", "age", 30)
slog.Error("failed to connect", "err", err)
```

输出：
```
2026-07-28T10:00:00.000+0800 INFO hello
2026-07-28T10:00:00.000+0800 INFO user logged in user=alice age=30
2026-07-28T10:00:00.000+0800 ERROR failed to connect err="connection refused"
```

### 自定义 logger

```go
opts := &slog.HandlerOptions{
    Level:     slog.LevelDebug,
    AddSource: true,  // 添加 source 位置
}
handler := slog.NewJSONHandler(os.Stderr, opts)
logger := slog.New(handler)
slog.SetDefault(logger)
```

### 分级

```go
slog.Debug("debug message")  // 默认不显示
slog.Info("info")
slog.Warn("warn")
slog.Error("error")

// 用 logger.InfoContext 携带 context
slog.InfoContext(ctx, "processing", "task", taskID)
```

### WithContext 加字段

```go
logger := slog.Default().With("request_id", reqID)
logger.Info("started")
logger.Info("done")
```

## 与现有项目的关系

`fing` 项目现在用自定义 logger（在 `log/logger.go`）。
可以**逐步迁移**到 slog，不影响现有 API。

```go
// 旧代码
log.Info("注册成功")
log.Errorf("失败:%v", err)

// 新代码（slog）
slog.Info("注册成功")
slog.Error("失败", "err", err)
```

## 我的理解

1. **slog 不强制替换**——它是新增包，旧 `log` 还在
2. **结构化字段是 key-value**——比 fmt.Sprintf 性能高
3. **Handler 可插拔**——可以写自己的 Handler 输出到任何地方
4. **JSON 输出适合生产**——文本输出适合开发

## 实战建议

- 新项目直接用 slog
- 旧项目保留 log 入口，逐步替换
- Handler 选型：
  - 开发：`slog.NewTextHandler` 人类可读
  - 生产：`slog.NewJSONHandler` 机器可解析
  - 自定义：写 `slog.Handler` 接口实现

## 参考

- 源码：`src/log/slog/`
- [Go 1.21 Release Notes](https://go.dev/doc/go1.21#enhanced_logging)
- 笔记时间：2026-07
