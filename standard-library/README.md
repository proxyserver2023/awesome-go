# Standard Library

## archive

### tar

```go
package main

import "fmt"

func main() {
	var buf bytes.Buffer
	tw := tar.NewWriter(&buf)
	var files = []struct{
		Name, Body string
	}{
		{"readme.txt", "this archive contains some text files."},
		{"adsfsadjkf.txt", "akfdljkaddfjl\nafadsfsad\n\asfdadsfsdafadsf\n"},
		{"adjfkajklfs.txt", "akdsjfkljfladf"},
	}

	for _, file := range files {
		hdr := &tar.Header{
			Name: file.Name,
			Mode: 0600,
			Size: int64(len(file.Body)),
		}

		if  err := tw.WriteHeader(hdr); err != nil {
			log.Fatal(err)
		}

		if _, err := tw.Write([]byte(file.Body)); err != nil {
			log.Fatal(err)
		}

	}

	if err := tw.Close(); err != nil {
		log.Fatal(err)
	}

	// Open and iterate through the files in the archive
	tr := tar.NewReader(&buf)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Contents of %s\n", hdr.Name)
		if _, err := io.Copy(os.Stdout, tr); err != nil {
			log.Fatal(err)
		}
		fmt.Println()
	}
}
```

## http

```go
resp, err := http.Get("http://example.com")
resp, err := http.Post("http://example.com/upload", "image/jpeg", &buf)
resp, err := http.PostForm(
                "http://example.com/form",
                url.Values{"key": {"Value"}, "id": {"123"}},
             )
```

the client must close the response body when finished with it:

```go
resp, err := http.Get("http://example.com")
if err != nil {
	// handle error
}

defer resp.Body.Close()
body, err := ioutil.ReadAll(resp.Body)
```

For control over HTTP client Headers, redirect-policy and other settings, create a Client:

```go
client := &http.Client{
	CheckRedirect: redirectPolicyFunc,
}

resp, err := client.Get("http://example.com")

// create a request for more flexibility
req, err := http.NewRequest("GET", "http://example.com", nil)
req.Header.Add("If-None-Match", `W/"wyzzy"`)
resp, err := client.Do(req)
```

### HTTP call with a HTTP Client

```go
request, reqErr := http.NewRequest("GET", "http://example.com", nil)
client := &http.Client{
        Timeout: time.Duration(10 * time.Second),
}
response, resErr := client.Do(request)
```

### Dump HTTP Request

```go
package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/gorilla/mux"
)

func DumpRequest(w http.ResponseWriter, req *http.Request) {
	requestDump, err := httputil.DumpRequest(req, true)
	if err != nil {
		fmt.Fprint(w, err.Error())
	} else {
		fmt.Fprint(w, string(requestDump))
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/get", DumpRequest).Methods("GET")
	router.HandleFunc("/post", DumpRequest).Methods("POST")
	log.Fatal(http.ListenAndServe(":12345", router))
}

```
