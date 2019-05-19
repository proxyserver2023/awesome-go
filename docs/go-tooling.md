# Go Tooling

```shell
$ cat > main.go
$ go run main.go
$ gofmt main.go # Outputs the formatted code in shell
$ gofmt -d main.go # Outputs the diff between formatted code and original code
$ gofmt -w main.go # Replace original with formatted code
$ go build
$ go install
$ go run main.go
$ go list -f '{{ .Name }}: {{ .Doc }}'
$ go list -f '{{ join .Imports "\n"}}'
$ go doc
$ go doc fmt
$ go doc fmt Printf
$ godoc -http :6060
```
