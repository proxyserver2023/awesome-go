# Concurrency

Do not communicate by sharing memory; instead, share memory by communicating.

## How GO's concurrency model is superior

Consider a program that polls a list of URLs.

### Traditional Data Sharing Technique

```go
type Resource struct {
    url        string
    polling    bool
    lastPolled int64
}

type Resources struct {
    data []*Resource
    lock *sync.Mutex
}

func Poller(res *Resources) {
    for {
        // get the least recently-polled Resource
        // and mark it as being polled
        res.lock.Lock()
        var r *Resource
        for _, v := range res.data {
            if v.polling {
                continue
            }
            if r == nil || v.lastPolled < r.lastPolled {
                r = v
            }
        }
        if r != nil {
            r.polling = true
        }
        res.lock.Unlock()
        if r == nil {
            continue
        }

        // poll the URL

        // update the Resource's polling and lastPolled
        res.lock.Lock()
        r.polling = false
        r.lastPolled = time.Nanoseconds()
        res.lock.Unlock()
    }
}
```

### GO's IDIOM for Data Sharing

```go
type Resource string

func Poller(in, out chan *Resource) {
        for r := range in {
                // Poll the URL
                // send the processed resource to out
                out <- r
        }
}
```

## Making Concurrent API Request

```
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type HTTPResponse struct {
	status string
	body   []byte
}

func DoHTTPGet(url string, ch chan<- HTTPResponse) {
	httpResponse, _ := http.Get(url)
	httpBody, _ := ioutil.ReadAll(httpResponse.Body)
	ch <- HTTPResponse{httpResponse.Status, httpBody}
}

func DoHTTPPost(url string, payload map[string]string, ch chan<- HTTPResponse) {
	jsonValue, _ := json.Marshal(payload)
	httpResponse, _ := http.Post(
		url,
		"application/json",
		bytes.NewBuffer(jsonValue),
	)
	httpBody, _ := ioutil.ReadAll(httpResponse.Body)
	ch <- HTTPResponse{httpResponse.Status, httpBody}
}

func main() {
	var ch chan HTTPResponse = make(chan HTTPResponse)
	urls := [2]string{
		"https://my-json-server.typicode.com/typicode/demo/posts",
		"https://my-json-server.typicode.com/typicode/demo/comments",
	}

	for _, url := range urls {
		go DoHTTPGet(url, ch)
	}

	// Get the response
	for range urls {
		// Use the response (<-ch).body
		fmt.Println((<-ch).status)
	}

	postUrls := "https://jsonplaceholder.typicode.com/posts"

	var myPostParam []map[string]string
	value1 := map[string]string{"title": "test1", "body": "body1", "userId": "1"}
	value2 := map[string]string{"title": "test2", "body": "body2", "userId": "2"}

	myPostParam = append(myPostParam, value1, value2)

	for _, value := range myPostParam {
		//For each URL call the DOHTTPPost function (notice the go keyword)
		go DoHTTPPost(postUrls, value, ch)
	}

	for range myPostParam {
		// Use the response (<-ch).body
		fmt.Println((<-ch).status)
	}

}
```

```go
package main

import (
    "fmt"
    "log"
    "net/http"
    "sync"
)

var urls = []string{
    "https://google.com",
    "https://tutorialedge.net",
    "https://twitter.com",
}

func fetch(url string, wg *sync.WaitGroup) (string, error) {
    resp, err := http.Get(url)
    if err != nil {
        fmt.Println(err)
        return "", err
    }
    wg.Done()
    fmt.Println(resp.Status)
    return resp.Status, nil
}

func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Println("HomePage Endpoint Hit")
    var wg sync.WaitGroup

    for _, url := range urls {
        wg.Add(1)
        go fetch(url, &wg)
    }

    wg.Wait()
    fmt.Println("Returning Response")
    fmt.Fprintf(w, "Responses")
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
    handleRequests()
}
```

## Share memory by communicating

```go
package main

import (
	"log"
	"net/http"
	"time"
)

const (
	numPollers     = 2                // number of Poller goroutines to launch
	pollInterval   = 60 * time.Second // how often to poll each URL
	statusInterval = 10 * time.Second // how often to log status to stdout
	errTimeout     = 10 * time.Second // back-off timeout on error
)

var urls = []string{
	"http://www.google.com/",
	"http://golang.org/",
	"http://blog.golang.org/",
}

// State represents the last-known state of a URL.
type State struct {
	url    string
	status string
}

// StateMonitor maintains a map that stores the state of the URLs being
// polled, and prints the current state every updateInterval nanoseconds.
// It returns a chan State to which resource state should be sent.
func StateMonitor(updateInterval time.Duration) chan<- State {
	updates := make(chan State)
	urlStatus := make(map[string]string)
	ticker := time.NewTicker(updateInterval)
	go func() {
		for {
			select {
			case <-ticker.C:
				logState(urlStatus)
			case s := <-updates:
				urlStatus[s.url] = s.status
			}
		}
	}()
	return updates
}

// logState prints a state map.
func logState(s map[string]string) {
	log.Println("Current state:")
	for k, v := range s {
		log.Printf(" %s %s", k, v)
	}
}

// Resource represents an HTTP URL to be polled by this program.
type Resource struct {
	url      string
	errCount int
}

// Poll executes an HTTP HEAD request for url
// and returns the HTTP status string or an error string.
func (r *Resource) Poll() string {
	resp, err := http.Head(r.url)
	if err != nil {
		log.Println("Error", r.url, err)
		r.errCount++
		return err.Error()
	}
	r.errCount = 0
	return resp.Status
}

// Sleep sleeps for an appropriate interval (dependent on error state)
// before sending the Resource to done.
func (r *Resource) Sleep(done chan<- *Resource) {
	time.Sleep(pollInterval + errTimeout*time.Duration(r.errCount))
	done <- r
}

func Poller(in <-chan *Resource, out chan<- *Resource, status chan<- State) {
	for r := range in {
		s := r.Poll()
		status <- State{r.url, s}
		out <- r
	}
}

func main() {
	// Create our input and output channels.
	pending, complete := make(chan *Resource), make(chan *Resource)

	// Launch the StateMonitor.
	status := StateMonitor(statusInterval)

	// Launch some Poller goroutines.
	for i := 0; i < numPollers; i++ {
		go Poller(pending, complete, status)
	}

	// Send some Resources to the pending queue.
	go func() {
		for _, url := range urls {
			pending <- &Resource{url: url}
		}
	}()

	for r := range complete {
		go r.Sleep(pending)
	}
}
```

## Block For Goroutine execution

```go
package main

import (
	"fmt"
	"sync"
)

func myFunc(waitgroup *sync.WaitGroup) {
	fmt.Println("Inside my goroutine")
	waitgroup.Done()
}

func main() {
	fmt.Println("Hello World")
	var waitgroup sync.WaitGroup
	waitgroup.Add(1)
	go myFunc(&waitgroup)
	waitgroup.Wait()
	fmt.Println("Finished Execution")
}

```
