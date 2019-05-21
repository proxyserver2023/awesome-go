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

```

## Share memory by communicating

```go
var urls = []string{
        "http://www.google.com/",
        "http://golang.org/",
        "http://blog.golang.org/",
}


// State represents the last known state of a URL.
type State struct {
        url string
        status string
}


```
