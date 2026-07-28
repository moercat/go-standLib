# Go 1.21+ 标准库新包阅读笔记 #1

> 阅读对象：`slices` · `maps` · `sort` 包（Go 1.21 引入正式 API）

## 背景

Go 1.18 之前没有泛型，标准库的工具函数都是 `[]int` / `[]string` 单独写。
Go 1.18 加入泛型，1.21 把社区常用的 `golang.org/x/exp/slices` / `exp/maps` 收编为正式包。

## slices 包关键函数

### 搜索

```go
slices.BinarySearch([]int{1, 3, 5, 7}, 5)  // → 2, true
slices.Contains([]int{1, 2, 3}, 2)         // → true
slices.Index([]int{1, 2, 3}, 2)           // → 1
slices.Min([]int{3, 1, 2})                // → 1
slices.Max([]int{3, 1, 2})                // → 3
```

### 排序

```go
slices.Sort([]int{3, 1, 2})                // → [1 2 3]
slices.SortFunc(people, func(a, b Person) int {
    return cmp.Compare(a.Age, b.Age)
})
slices.IsSorted([]int{1, 2, 3})           // → true
slices.Reverse([]int{1, 2, 3})            // → [3 2 1]
```

### 复制与替换

```go
slices.Clone([]int{1, 2, 3})              // 浅拷贝
slices.Replace([]int{1, 2, 3, 4}, 1, 3, []int{9, 9})  // [1 9 9 4]
slices.Insert([]int{1, 3}, 1, 2)         // [1 2 3]
slices.Delete([]int{1, 2, 3}, 1, 2)      // [1 3]
slices.Compact([]int{1, 1, 2, 2, 3})     // [1 2 3]
```

### 批量操作

```go
slices.Equal([]int{1, 2}, []int{1, 2})   // true
slices.Compare([]int{1, 2}, []int{1, 3}) // -1
```

## maps 包关键函数

```go
m := map[string]int{"a": 1, "b": 2}
maps.Keys(m)    // → []string{"a", "b"}  (乱序)
maps.Values(m)  // → []int{1, 2}          (乱序)

maps.Clone(m)                              // 浅拷贝
maps.Copy(dst, src)                        // 复制 src 到 dst
maps.DeleteFunc(m, func(k string, v int) bool {
    return v < 0
})
maps.Equal(map[string]int{"a": 1}, map[string]int{"a": 1})  // true
```

## sort 包变化

`sort.Ints` / `sort.Strings` 标记为 deprecated，推荐用 `slices.Sort`。

```go
// 1.21 之前
sort.Ints([]int{3, 1, 2})
sort.Slice(people, func(i, j int) bool { return people[i].Age < people[j].Age })

// 1.21 之后
slices.Sort([]int{3, 1, 2})
slices.SortFunc(people, func(a, b Person) int { return cmp.Compare(a.Age, b.Age) })
```

## 自己实现一遍（学习用）

参考 `collections/` 目录下的简化重写。

## 我的理解

1. **`slices` 是 1.21 的"补完"**——它不创新，只是把社区的 x/exp 收编
2. **`sort` 被 slices 接管**—— `sort.Slice` 现在被视为旧 API
3. **没有 `Filter` / `Map` / `Reduce`**——slices 包有意保持"小而精"，复杂操作留给社区（oper / lo / go-funk）

## 参考

- [Go 1.21 Release Notes](https://go.dev/doc/go1.21)
- 源码：`src/slices/slices.go` `src/maps/maps.go`
- 笔记时间：2026-07
