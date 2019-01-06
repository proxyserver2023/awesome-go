# Awesome Golang

## Tour Of Golang
* Multiple Results

``` go
func swap(x, y string) (string, string) {
    return y, x
}

func main() {
    a, b := swap("hello", "world")
    fmt.Println(a, b)
}
```
* Named Return Values

``` go
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return
}
```

* Stacking Defers

``` go
func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
```

* Pointers

``` go
primes := [6]int{1,2,3,4,5,6}
var s []int = primes[1:4]
```

* Slices are like references to arrays

``` go
names := [4]string{
    "John", "Paul", "George", "Ringo",
}
fmt.Println(names)

a := names[0:2]
b := names[1:3]
fmt.Println(a, b)

b[0] = "XXX"
fmt.Println(a, b)
fmt.Println(names)

/*
Outputs
-------
[John Paul George Ringo]
[John Paul] [Paul George]
[John XXX] [XXX George]
[John XXX George Ringo]
*/
```
* Slice Literals

``` go
q := []int{2, 3, 5, 7, 11, 13}
fmt.Println(q)

r := []bool{true, false, true, true, false}
fmt.Println(r)

s := []struct{
        i int
        b bool
    }{
            {2, true},
            {1, false},
            {3, true},
            {1, true},
    }
```
First an array is created then builds a slice that references it.

* Slice Length and capacity
  + length: number of element it contains
  + capacity: number of elements in the underlying array, counting fromt the first element in the slice.
  + we can extend slice if it has enough capacity. extending beyond capacity throws `panic: runtime error: slice bounds out of range`


``` go
s := []int{1,2,3,4,5,6,7}
// len(s)     => 7
// cap(s)     => 7
// Println(s) => [1,2,3,4,5,6,7]

s = s[:0]
// len(s)     => 0
// cap(s)     => 7
// Println(s) => []

s = s[:2] // extending it 2nd element
// len(s)     => 2
// cap(s)     => 7
// Println(s) => [1 2]

s = s[2:] // droping its first two values
// len(s)     => 0
// cap(s)     => 5
// Println(s) => []
```

* Nil value of slice
  - the zero value of slice is nil
  - nil slice has a length and capacity of 0 and has `no underlying array`.

```go
var s []int
// Println(s) => []
// len(s)     => 0
// cap(s)     => 0

if s == nil {
    fmt.Println("nil!")
}
```

* Creating a slice with make
the `make` function allocates a zeroed array and returns a slice that refers to the array.

``` go
a := make([]int, 5) // len(a) = 5
```
to specify a capacity, pass a third arg
``` go
b := make([]int, 0, 5) // len(b) = 0; cap(b) = 5
b = b[:cap(b)]         // len(b) = 5; cap(b) = 5
b = b[1:]              // len(b) = 4; cap(b) = 4
```

* Slices of slices

``` go
board := [][]string{
    []string{"_", "_", "_"},
    []string{"_", "_", "_"},
    []string{"_", "_", "_"},
}
```

* Appending to a slice

``` go
var s, t []int
t = []int{4,5,6}
s = append(s, 1)
s = append(s, 2)
s = append(s, 3)
s = append(s, t...)
fmt.Println(s) => [1,2,3,4,5,6]
```

* Exercise: Slices

``` go
package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	s := make([][]uint8, 0)
	for i := 0; i < dy; i++ {
		t := make([]uint8, 0)
		for j := 0; j < dx; j++ {
			n := ((i*j + i^j) + (i*j + i^j))/2
			t = append(t, uint8(n))
		}
		s = append(s, t)
	}
	return s
}

func main() {
	pic.Show(Pic)
}

```
