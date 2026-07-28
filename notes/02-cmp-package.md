# Go 1.21+ 标准库新包阅读笔记 #2

> 阅读对象：`cmp` · `constraints` · `slices` 内部使用的 `cmp.Ordered`

## cmp 包（Go 1.21）

只有 4 个函数，但用途极广。

```go
import "cmp"

// Compare 返回 -1/0/1
cmp.Compare(1, 2)          // -1
cmp.Compare(2, 2)          // 0
cmp.Compare(3, 2)          // 1
cmp.Compare("a", "b")      // -1

// Less 简化版
cmp.Less(1, 2)             // true

// Ordereed 是 1.21 新引入的类型约束
// 替代 1.18 时的 `interface{ ~int | ~int32 | ... | ~string }`
func Min[T cmp.Ordered](a, b T) T {
    if a < b { return a }
    return b
}

// Compare 用作 sort.Less 替代
slices.SortFunc(people, func(a, b Person) int {
    return cmp.Compare(a.Age, b.Age)
})
```

## 自带 `min` / `max` 内建函数（Go 1.21）

```go
min(1, 2, 3)              // 1
max(1, 2, 3)              // 3
min("a", "b")             // "a"
```

**注意**：内建 `min`/`max` 只接受有序类型（`cmp.Ordered`），不像泛型 `Min` 可以传 `cmp.Compare` 函数。

## 实际场景

### 排序自定义结构

```go
type User struct {
    Name string
    Age  int
}

users := []User{{"alice", 30}, {"bob", 25}}

// 用 cmp.Compare 一行搞定
slices.SortFunc(users, func(a, b User) int {
    return cmp.Compare(a.Age, b.Age)
})
```

### 链式比较

```go
// 先按 Age 升序，相同 Age 按 Name 升序
slices.SortFunc(users, func(a, b User) int {
    if c := cmp.Compare(a.Age, b.Age); c != 0 {
        return c
    }
    return cmp.Compare(a.Name, b.Name)
})
```

## 我的理解

- `cmp.Compare` 是 1.21 的"类型安全 compare"——以前写 `if a < b` 在泛型里要写 `interface{...}`
- 自带 `min/max` 让简单场景不用 import 包
- 但 Go 没引入 `<=>` 操作符（即使很多语言有），所以 `cmp.Compare` 还是会成为标配

## 参考

- 源码：`src/cmp/cmp.go`
- 笔记时间：2026-07
