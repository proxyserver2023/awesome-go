# Graceful Shutdown

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
