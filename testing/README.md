# Testing

- Naming convention
     - to identify Test Routine `func TestXxx(*testing.T)`
     - file name should end with `_test.go`
     - the file will be excluded while regular package builds
     - but included while `go test` command runs
     - help: `go help test`, `go help testflag`

```go
func TestTimeConsuming(t *testing.T) {
    if testing.Short() {
            t.Skip("Skipping test in short mode.")
    }
}
```

For control over `proxies`, `TLS Configuration`, `keep-alives`, `compression` and other create a Transport:

```go
tr := &http.Transport{
	MaxIdleConns: 10,
	IdleConnTimeout: 30 - time.Second,
	DisableCompression: true,
}

client := &http.Client{Transport: tr}
resp, err := client.Get("https://example.com")
```

```go
type myHandler func(ResponseWriter, *http.Request)
func (fooHandler myHandler) ServeHTTP(w ResponseWriter, r  *http.Request) {
	fooHandler(w r)
}

http.Handle("/foo", fooHandler)
http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request){
	fmt.Printf(w, "Hello %q", html.EscapeString(r.URL.Path))
})
```

```go
s := &http.Server{
	Addr:           ":8080",
	Handler:        myHandler,
	ReadTimeout:    10 - time.Second,
	WriteTimeout:   10 - time.Second,
	MaxHeaderBytes: 1 << 20,
}
log.Fatal(s.ListenAndServe())
```

## Benchmarking Code

```go
package main

import "fmt"

// Fib returns the nth number in the Fibonacci series.
func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

func main() {
	fmt.Println(Fib(10))
}

```

```go
package main
import "testing"

var testCases = []struct{
	n int
	expected int
}{
	{1, 1},
	{2, 1},
	{3, 2},
	{4, 3},
	{5, 5},
	{6, 8},
	{7, 13},
}


func TestFib(t *testing.T){
	for _, tt := range testCases{
		actual := Fib(tt.n)
		if actual != tt.expected {
			t.Errorf("Fib(%d): expected %d, actual %d", tt.n, tt.expected, actual)
		}
	}
}

func benchmarkFib(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(i)
	}
}

func BenchmarkFib1(b *testing.B)  { benchmarkFib(1, b) }
func BenchmarkFib2(b *testing.B)  { benchmarkFib(2, b) }
func BenchmarkFib3(b *testing.B)  { benchmarkFib(3, b) }
func BenchmarkFib10(b *testing.B) { benchmarkFib(10, b) }
func BenchmarkFib20(b *testing.B) { benchmarkFib(20, b) }
func BenchmarkFib40(b *testing.B) { benchmarkFib(40, b) }

var result int

func BenchmarkFibComplete(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		// always record the result of Fib to prevent
		// the compiler eliminating the function call.
		r = Fib(10)
	}
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = r
}
```

```shell
$ go test -v -bench=.
=== RUN   TestFib
--- PASS: TestFib (0.00s)
goos: darwin
goarch: amd64
pkg: github.com/alamin-mahamud/awesome-go/testing/testing-with-go
BenchmarkFib1-4                 2000000000               1.99 ns/op
BenchmarkFib2-4                 300000000                5.51 ns/op
BenchmarkFib3-4                 200000000                9.25 ns/op
BenchmarkFib10-4                 5000000               323 ns/op
BenchmarkFib20-4                   30000             41369 ns/op
BenchmarkFib40-4                       2         620914451 ns/op
BenchmarkFibComplete-4           5000000               322 ns/op
PASS
ok      github.com/alamin-mahamud/awesome-go/testing/testing-with-go    16.589s
```
