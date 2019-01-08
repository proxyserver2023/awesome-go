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

* Maps

``` go
type Vertex struct{Lat, Long float64}
var m map[string]Vertex

func main() {
    m = make(map[string]Vertex)
    m["Bell Labs"] = Vertex{40.12312, -12.12312}
    fmt.Println(m["Bell Labs"])
}
```

* Map literals

``` go
type Vertex struct {Lat, Long float64}

var m = map[string]Vertex{
    "Bell Labs": Vertex{40.6833, -74.39976},
    "Google": Vertex{37.123, -12.12},
}
```

* Mutating Maps

``` go
m := make(map[string]interface{})
m["alamin"] = "Mahamud"
name = m["alamin"]
delete(m, "alamin")
element, ok := m["alamin"]
```

* Exercise Maps

``` go
func WordCount(s string) map[string]int {
	m := make(map[string]int)
	words := strings.Fields(s)
	for _, word := range words {
		m[word]++
	}
	return m
}
```
* Functions as values

``` go
func compute(fn func(float64, float64) float64) float64 {
    return fn(3, 4)
}


fun main() {
    hypot := func(x, y float64) float64 {
            return math.Sqrt(x*x + y*y)
    }
    fmt.Println(compute(hypot))
    fmt.Println(compute(math.Pow))
}
```

* Function closures
``` go
func adder() func (int) int {
    sum := 0
    return func (x int) int {
            sum += x
            return sum
    }
}

func main() {
    pos, neg := adder(), adder()
    for i := 0; i < 100; i++ {
            fmt.Println(
                    pos(i),
                    neg(-2*i),
            )
    }
}
```

* Exercise Fibonacci Series

``` go
func fib() func() int {
    n := 0
    a := 0
    b := 1
    return func() int {
            var ret int
            switch {
                    case n == 0:
                            n++
                            ret = 0
                    case n == 1:
                            n ++
                            ret = 1
                    default:
                            ret = a + b
                            a, b = b, a+b
            }
            return ret
    }
}

func main() {
    f := fib()
    for i := 0; i < 10; i++ {
            fmt.Println(f())
    }
}
```
* Choosing a value or pointer receiver
  - if you use pointer receiver data is refernced, not copied so you can modify it and don't need to use storage
``` go
type Vertex struct{
    X, Y float64
}

func (v *Vertex) Scale(f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

```
* Interface values
  - `(value, type)`
  - an interface value holds a value of a specific underlying concrete type.
  - calling a method on an interface value executes the method of the same name on its underlying type.

``` go
type I interface{
    M()
}

type T struct{
    S string
}

func (t *T) M() {fmt.Println(t.S)}

type F float64
func (f *F) M() {fmt.Println(f)}

```
* Interface values with nil underlying values
  - If the concrete value inside the interface itself is nil, the method will be called with a nil receiver.

``` go
type I interface{M()}
type T struct{S string}

func (t *T) M() {
    if t == nil {
            fmt.Println("<nil>")
            return
    }
    fmt.Println(t.S)
}

func describe(i I){
    fmt.Printf("(%v, %T)", i, i)
}

func main() {
    var i I
    var t *T
    i = t; describe(i); i.M()
    i = &T{"Hello"}; describe(i); i.M()
}

/*
Outputs
-----------
(<nil>, *main.T)
<nil>
(&{hello}, *main.T)
hello

*/

```
* Nil interface values
  - A nil interface value holds neither value nor concrete type.
  - Calling a method on a nil interface is a run-time error because there is no type inside the interface tuple to indicate which concrete method to call.

``` go
type I interface {
	M()
}

func main() {
	var i I
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

/*
Outputs
---------------
(<nil>, <nil>)
panic: runtime error: invalid memory address or nil pointer dereference

*/

```
* The empty interface
  - The interface type that specifies zero methods is known as the empty interface
  - `interface{}`
  - An empty interface may hold values of any type. (Every type implements at least zero methods.)
  - Empty interfaces are used by code that handles values of unknown type. For example, fmt.Print takes any number of arguments of type `interface{}`.
``` go
func main() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

/*
    Outputs
    --------
    (<nil>, <nil>)
    (42, int)
    ("hello", string)
*/

```

* Type assertions
  - `t := i.(T)`
  - a type assertion provides access to an interface value's underlying concrete value.
  - `i` holds the concrete type `T` and assigns the underlying `T` value to the variable `t`.
  - if `i` does not hold `T`, the statement will trigger a `panic`.
``` go
var i interface{} = "Hello"

s := i.(string)
fmt.Println(s) // ---> hello

s, ok :=  i.(string)
fmt.Println(s, ok) // ---> hello, true

f, ok := i.(float64)
fmt.Println(f, ok) // ---> 0, false & panic

```
* Type Switches

``` go
switch v := i.(type){
    case T:
    // here v has type T
    case S:
    // here v has type S
    default:
    // no match; here v has the same type as i
}
```

``` go
func do(i interface{}) {
    switch v := i.(type) {
    case int:
            fmt.Printf("Twice %v is %v\n", v, v*2)
    case string:
            fmt.Printf("%q is %v bytes long\n", v, len(v))
    default:
            fmt.Printf("I don't know about type %T!\n", v)
    }
}

func main() {
    do(21)
    do("hello")
    do(true)
}

/*

Outputs
-------

Twice 21 is 42
"hello" is 5 bytes long
I don't know about type bool!

*/
```
* Stringers

``` go
type Stringer interface{
    String() string
}

type Person struct {
    Name string
    Age int
}

func (p Person) String() string{
    return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}
```

* Exercise Stringers

``` go
package main

import "fmt"

type IPAddr [4]byte

func (i IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", i[0], i[1], i[2], i[3])
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

/*

Outputs
--------

loopback: 127.0.0.1
googleDNS: 8.8.8.8

*/

```

* Errors

``` go
type error interface {
    Error() string
}
```

``` go
type MyError struct {
    When time.Time
    What string
}

func(e *MyError) Error() string{
    return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

/* Outputs
--------------
at 2009-11-10 23:00:00 +0000 UTC m=+0.000000001, it didn't work
*/
```

* Exercise errors

``` go
package main

import (
	"fmt"
)


type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		e := ErrNegativeSqrt(x)
		return 0, e
	}
	return 0, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

```
* Reader

``` go
package main

import (
    "fmt"
    "io"
    "strings"
)

func main() {
    r := strings.NewReader("Hello, Reader!")
    b := make([]byte, 8)

    // loop until EOF to read from the Reader
    for {
            n, err := r.Read(b)
            fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
            fmt.Printf("b[:n] = %q\n", b[:n])

            // Breaking Condition
            if err == io.EOF {
                    break
            }
    }
}

/*
Outputs
----------
n = 8 err = <nil> b = [72 101 108 108 111 44 32 82]
b[:n] = "Hello, R"
n = 6 err = <nil> b = [101 97 100 101 114 33 32 82]
b[:n] = "eader!"
n = 0 err = EOF b = [101 97 100 101 114 33 32 82]
b[:n] = ""
*/
```

* Exercise Reader

``` go
package main

import (
    "fmt"
    "io"
    "os"
)

type MyReader struct{}

func (m MyReader) Read(b []byte) (int, error) {
	b[0] = 'A'
	return 1, nil
}

func main() {
	Validate(MyReader{})
}

func Validate(r io.Reader) {
    b :=  make([]byte, 1024, 2048)
    i, o := 0, 0

    for ; i < 1<<20 && o < 1<<20; i++ {
            n, err := r.Read(b)

            for i, v := range b[:n] {
                    if v != 'A' {
                            fmt.Fprintf(os.Stderr, "got byte %x at offset %v, want 'A'\n", v, o+i)
                            return
                    }
            }

            o += n
            if err != nil {
                    fmt.Fprintf(os.Stderr, "read error: %v\n", err)
                    return
            }

    }

    if o == 0 {
            fmt.Fprintf(os.Stderr, "read zero bytes after %d Read Calls\n", i)
    }
    return
}
fmt.Println("OK!")

```
* Exercise Custom Reader

``` go
package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13Sub(l *byte) {
	// check small character
	if *l >= 97 && *l <= 122 {
		if *l + 13 > 122 {
			p := 122 - *l
			*l = 97 + p
		} else {
			*l = *l + 13
		}

	} else if *l >= 65 && *l <= 90 {
		if *l + 13 > 90 {
			p := 90 - *l
			*l = 65 + p
		} else {
			*l = *l + 13
		}
	}

}

func (r13 rot13Reader) Read(b []byte) (int, error) {
	n, err := r13.r.Read(b)
	for i, _ := range b[:n] {
		rot13Sub(&b[i])
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

```
* Images

``` go
import (
	"fmt"
	"image"
)

func main() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}

```
