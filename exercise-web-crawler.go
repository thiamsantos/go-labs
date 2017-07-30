package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

type Result struct {
	url, body string
	depth     int
}

type UrlCache struct {
	v   map[string]bool
	mux sync.Mutex
}

func (cache *UrlCache) Has(url string) bool {
	cache.mux.Lock()
	_, ok := cache.v[url]
	cache.mux.Unlock()
	return ok
}

func (cache *UrlCache) Set(url string) {
	cache.mux.Lock()
	cache.v[url] = true
	cache.mux.Unlock()
}

func Crawl(url string, depth int, fetcher Fetcher, ch chan Result, cache UrlCache) {
	defer close(ch)

	if depth < 0 || cache.Has(url) {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	cache.Set(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	ch <- Result{url, body, depth}

	for _, url := range urls {
		recursiveChannel := make(chan Result)
		go Crawl(url, depth-1, fetcher, recursiveChannel, cache)

		for result := range recursiveChannel {
			ch <- result
		}
	}
	return
}

func main() {
	ch := make(chan Result)
	cache := UrlCache{v: make(map[string]bool)}
	go Crawl("http://golang.org/", 4, fetcher, ch, cache)

	for value := range ch {
		fmt.Printf("found: %v\n", value)
	}
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
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
