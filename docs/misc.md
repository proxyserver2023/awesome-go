# Misc

## `defer *http.Request.Body.Close()`

A request body does not need to be closed in the handler. From the http.Request documentation

```go
// The Server will close the request body. The ServeHTTP
// Handler does not need to.
```

## compare two slices

- Basic Case

```go
// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func Equal(a, b []int) bool {
    if len(a) != len(b) {
        return false
    }
    for i, v := range a {
        if v != b[i] {
            return false
        }
    }
    return true
}
```

- Optimized code for byte slices

```go
// To compare byte slices, use the optimized bytes.Equal. This function also treats nil arguments as equivalent to empty slices.
```

- General Purpose code for recursive comparison

For testing purposes, you may want to use reflect.DeepEqual. It compares two elements of any type recursively.

```go
var a []int = nil
var b []int = make([]int, 0)
fmt.Println(reflect.DeepEqual(a, b)) // false
```

The performance of this function is much worse than for the code above, but it’s useful in test cases where simplicity and correctness are crucial. The semantics, however, are quite complicated.
