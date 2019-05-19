# Slices

Arrays examples

```go
var a [4]int
a[0] = 1
i := a[0]
// i == 1
```

Slice Syntax

```go
func make([]T, len, cap) []T
```

```go
s := make([]byte, 5) // Capacity == Length by default
```
