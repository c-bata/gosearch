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
func Crawl(url string, depth int, fetcher Fetcher) {
	type URL struct {
		url   string
		depth int
	}

	msg := make(chan string)
	tocrawl := make(chan URL)
	quit := make(chan int)

	crawler := func(url string, depth int) {
		defer func() { quit <- 0 }()

		if depth <= 0 {
			return
		}

		body, urls, err := fetcher.Fetch(url)

		if err != nil {
			msg <- fmt.Sprintf("%s\n", err)
			return
		}

		msg <- fmt.Sprintf("found: %s %q\n", url, body)

		for _, u := range urls {
			tocrawl <- URL{u, depth - 1}
		}
	}

	works := 1

	crawled := make(map[string]bool)
	crawled[url] = true

	go crawler(url, depth)

	for works > 0 {
		select {
		case s := <-msg:
			fmt.Printf(s)
		case r := <-tocrawl:
			if !crawled[r.url] {
				crawled[r.url] = true
				works++

				go crawler(r.url, r.depth)
			}
		case <-quit:
			works--
		}
	}
}

func main() {
	var seed string = "http://golang.org/"
	Crawl(seed, 4, fetcher)
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
