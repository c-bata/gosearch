package crawler

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

type URL struct {
	Url   string
	Depth int
}

type CrawlResponse struct {
	StatusCode int
	Url        string
	Body       string
}

func getAllLinks(body string) (urls []string) {
	r := regexp.MustCompile(`https?://[\w/:%#\$&\?\(\)~\.=\+\-]+`)
	urls = r.FindAllString(body, -1)
	return urls
}

func fetch(url string, depth int, resp chan CrawlResponse, tocrawl chan URL) {
	if depth <= 0 {
		return
	}

	r, err := http.Get(url)
	if err != nil {
		return
	} else {
		t := r.Header.Get("Content-Type")
		if t == "text/css" || t == "text/javascript" || strings.HasPrefix(t, "image") {
			return
		}
	}
	defer r.Body.Close()
	bytesBody, err := ioutil.ReadAll(r.Body)
	body := string(bytesBody)
	if err != nil {
		return
	}
	func() {
		resp <- CrawlResponse{
			StatusCode: r.StatusCode,
			Url:        url,
			Body:       body,
		}
	}()

	for _, url := range getAllLinks(body) {
		tocrawl <- URL{Url: url, Depth: depth - 1}
	}
	return
}

func Crawl(seed string, depth int, resp chan CrawlResponse) {
	tocrawl := make(chan URL)
	crawled := make(map[string]bool) // TODO: save hashed value

	crawled[seed] = true
	go fetch(seed, depth, resp, tocrawl)

	for u := range tocrawl {
		if !crawled[u.Url] {
			crawled[u.Url] = true
			go fetch(u.Url, u.Depth, resp, tocrawl)
		}
	}
	close(resp)
}
