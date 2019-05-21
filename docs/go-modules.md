# Go Modules

## Requirements

- `go - 1.11+`

## Create a Hello World Application using go mod

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
