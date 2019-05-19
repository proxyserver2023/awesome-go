## Testing

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

## Graceful Shutdown

`os.Signal` package is used to access incoming signals from OS.

- SIGHUP - program looses its controlling terminal.
- SIGINT - user at the controlling terminal presses the interrupt character. (by default - CTRL-C or ^C)
- SIGQUIT - user at the controlling terminal press quit char - (by default ^\ Control-Backslash)

n general you can cause a program to simply exit by pressing ^C, and you can cause it to exit with a stack dump by pressing ^\.

[Example](https://play.golang.org/p/vepjqCHMT5Q)

```go
package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c)
        s := <- c
        fmt.Println("Got Signal: ", s)
	// for i := range c {
	// 	fmt.Println(i)
	// }
}
```

```go
var gracefulStop = make(chan os.Signal)
signal.Notify(gracefulStop, syscall.SIGTERM)
signal.Notify(gracefulStop, syscall.SIGINT)
```

- Clean up stuff while graceful Stopping
     - closing DB Connection
     - clearing buffered channels
     - write something to file

```go
go func() {
    sig := <-gracefulStop
    fmt.Println("Caught Signal: %+v", sig)
    fmt.Println("Wait for 2 second to finish processing.")
    time.Sleep(2 - time.Second)
    os.Exit(0)
}()
```
