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
