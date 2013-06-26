package main

import (
    "fmt"
)

type Fetcher interface {
    // Fetch returns the body of URL and
    // a slice of URLs found on that page.
    Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func grab(url string, depth int, fetcher Fetcher, ch chan urlD,ch2 chan int) {
    body, urls, err := fetcher.Fetch(url)   
    if err != nil {
        fmt.Println(err)
        return
    }        
    fmt.Printf("found: %s %q\n", url, body)
    for _, u := range urls {
        result:=urlD{u,depth}
        ch <- result
    }
    ch2 <- 0
    return
}

type urlD struct {
	url string
    depth int
}

func Crawl(url string, depth int, fetcher Fetcher){
    m:=make(map[string]bool)
    ch:=make(chan urlD)
    ch2:=make(chan int)
    count:=0    
    go func(){ 
        ch <- urlD{url,depth}
    }()
    
    for {
        select {
        case v := <- ch:
            if m[v.url]!=true && v.depth >0 {
                m[v.url]=true
                count++
                go grab(v.url, v.depth-1, fetcher, ch,ch2)
            }
        case <- ch2:
            count--
            if count == 1 {
            	return
            }
            
        }

    }
}

func main() {
    Crawl("http://golang.org/", 4, fetcher);
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
    body string
    urls []string
}

func (f *fakeFetcher) Fetch(url string) (string, []string, error) {
    if res, ok := (*f)[url]; ok {
        return res.body, res.urls, nil
    }
    return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = &fakeFetcher{
    "http://golang.org/": &fakeResult{
        "The Go Programming Language",
        []string{
            "http://golang.org/pkg/",
            "http://golang.org/cmd/",
        },
    },
    "http://golang.org/pkg/": &fakeResult{
        "Packages",
        []string{
            "http://golang.org/",
            "http://golang.org/cmd/",
            "http://golang.org/pkg/fmt/",
            "http://golang.org/pkg/os/",
        },
    },
    "http://golang.org/pkg/fmt/": &fakeResult{
        "Package fmt",
        []string{
            "http://golang.org/",
            "http://golang.org/pkg/",
        },
    },
    "http://golang.org/pkg/os/": &fakeResult{
        "Package os",
        []string{
            "http://golang.org/",
            "http://golang.org/pkg/",
        },
    },
}
