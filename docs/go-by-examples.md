# Go By Examples

- Hello World

```go
package main
import "fmt"
func main() {
    fmt.Println("hello world")
}
```

```bash
go run hello-world.go
go build hello-world.go
./hello-world
```

- Values

```go
package main
import "fmt"
func main() {
    // Strings, which can be added together with +

    fmt.Println("go" + "lang")

    fmt.Println("1+1 =", 1+1)
    fmt.Println("7.0/3.0 =", 7.0/3.0)

    fmt.Println(true && false)
    fmt.Println(true || false)
    fmt.Println(!true)
}
```

```bash
$ go run values.go
golang
1+1 = 2
7.0/3.0 = 2.3333333333333335
false
true
false
```

- Go Variables

```go
var a = "initial"
fmt.Println(a)

var b, c int = 1, 2
fmt.Pritnln(b, c)

var d = true
fmt.Println(d)

var e int
fmt.Println(e) // Zero Value

f := "short"
fmt.Println(f)
```

```bash
$ go run variables.go
initial
1 2
true
0
short
```

- Constants
     - A numeric constant has no type until it’s given one, such as by an explicit cast.
     - A number can be given a type by using it in a context that requires one, such as a variable assignment or function call. For example, here math.Sin expects a float64.

```go
const s string = "constant"

fmt.Println(s)
const n = 50000000
const d = 3e20/n
fmt.Println(int64(d))
fmt.Println(math.Sin(n))
```

```bash
$ go run constant.go
constant
6e+11
600000000000
-0.28470407323754404
```

- For

```go
i := 1
// Single condition
for i <= 3 {
    fmt.Println(i)
    i = i + 1
}

// A classic initial/condition/after for loop.

for j := 7; j <= 9; j++ {
    fmt.Println(j)
}

// for without a condition will loop repeatedly until you break out of the loop or return from the enclosing function.

for {
    fmt.Println("loop")
    break
}

// You can also continue to the next iteration of the loop.
for n := 0; n <= 5; n++ {
    if n%2 == 0 {
        continue
    }
    fmt.Println(n)
}
```

- If else

```go
if 7%2 == 0 {
    fmt.Println("7 is even")
} else {
    fmt.Println("7 is odd")
}

// You can have an if statement without an else.

if 8%4 == 0 {
    fmt.Println("8 is divisible by 4")
}

// A statement can precede conditionals; any variables declared in this statement are available in all branches.

if num := 9; num < 0 {
    fmt.Println(num, "is negative")
} else if num < 10 {
    fmt.Println(num, "has 1 digit")
} else {
    fmt.Println(num, "has multiple digits")
}
```

- No ternary operator in golang

- Switch

```go
i := 2
fmt.Println("Write ", i, " as ")
switch i {
case 1:
    fmt.Println("one")
case 2:
    fmt.Println("two")
case 3:
    fmt.Println("three")
}

switch time.Now().Weekday() {
case time.Saturday, time.Sunday:
    fmt.Println("It's the weekend")
default:
    fmt.Println("It's a weekday")
}

t := time.Now()
switch {
case t.Hour() < 12:
    fmt.Println("It's before noon.")
default:
    fmt.Println("It's after noon.")
}

whatAmI := func(i interface{}) {
    switch t := i.(type)
    case bool:
            fmt.Println("I'm a bool")
    case int:
            fmt.Println("I'm an int")
    default:
            fmt.Printf("Don't know the type%T\n", t)
}

whatAmI(true)
whatAmI(1)
whatAmI("hey")
```

## Go Modules

### Requirements

- `go - 1.11+`

### Create a Hello World Application using go mod

```bash
mkdir -p /tmp/hello
go mod init github.com/alamin-mahamud/go-hello
touch hello.go
```

hello.go

```go
package main

import (
	"fmt"
	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Hello())
}
```

```bash
$ go build
$ ./hello
Hello, world
$ go list -m all
$ LANG=fr ./hello
Bonjour le monde
$ go list -m -u all
$ go get -u golang.org/x/text
```

- test all packages with standard library

```bash
$ go test -short all
```

- test a certain package

```bash
go test rsc.io/quote/...
```

- upgrade all modules

```bash
go get -u
```

- downgrading deps

```bash
$ go list -m -versions rsc.io/sampler
rsc.io/sampler v1.0.0 v1.2.0 v1.2.1 v1.3.0 v1.3.1 v1.99.99
$ go get rsc.io/sampler@v1.2.0
$ go list -m all
```

- using a fork package

```bash
$ git clone <fork-repo-url> <location>
$ fork_repo_location=location
$ cd $fork_repo_location
$ versionNumber=$(cd /your-package-location && go list -m -f "{{.Version}}" $package_url)
$ git checkout -b new-branch-name $versionNumber
```

```bash
# in your own package
$ go mod edit -replace "pacakge_url=$fork_repo_location"
```

- using a updated remote package

```bash
# cd into forked package
git remote set-url origin <your_url>
git add .
git commit -m "Some Changes"
git push origin <branch_name>

tagName=v1.2.3-rc1
git tag $tagName
git push origin $tagName

# cd into your project
originalPackageUrl="rsc.io/quote"
updatePackageUrl="github.com/golab-bd/quote"
go mod edit -replace "$originPackageUrl=$updatePackageUrl@$tagName"
go list -m all
go build
```
