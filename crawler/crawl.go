package crawler

import (
	"io/ioutil"
	"net/http"
	"regexp"
)

type URL struct {
	url   string
	depth int
}

func getAllLinks(body string) (urls []string) {
	r := regexp.MustCompile(`https?://[\w/:%#\$&\?\(\)~\.=\+\-]+`)
	urls = r.FindAllString(body, -1)
	return urls
}

func crawl(url string, depth int, msg chan string, tocrawl chan URL) {
	if depth <= 0 {
		return
	}
	defer func() { msg <- url + " is crawled." }()

	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	for _, url := range getAllLinks(string(body)) {
		tocrawl <- URL{url: url, depth: depth - 1}
	}
	return
}

func Crawl(seed string, depth int, msg chan string) {
	tocrawl := make(chan URL)
	crawled := make(map[string]bool)

	crawled[seed] = true
	go crawl(seed, depth, msg, tocrawl)

	for u := range tocrawl {
		if !crawled[u.url] {
			crawled[u.url] = true
			go crawl(u.url, u.depth, msg, tocrawl)
		}
	}
	close(msg)
}
