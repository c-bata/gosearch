package crawler

import (
	"fmt"
	"github.com/c-bata/gosearch/env"
	"github.com/c-bata/gosearch/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllLinks(t *testing.T) {
	assert := assert.New(t)
	var input string = "hoge<a href=\"http://example.com/\">link title</a>hoge\n<a href=\"http://hoge.example.com/\">link title</a>"
	urls := getAllLinks(input)
	assert.Equal("http://example.com/", urls[0])
	assert.Equal("http://hoge.example.com/", urls[1])
}

func DummyCrawledHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hogehoge<a href=\"http://example.com/\">fugafuga")
}

func TestFetch(t *testing.T) {
	assert := assert.New(t)
	env.Init()
	err := models.Dialdb(env.GetDBHost())
	assert.Nil(err)
	defer models.Session.Close()

	ts := httptest.NewServer(http.HandlerFunc(DummyCrawledHandler))
	defer ts.Close()

	resp := make(chan CrawlResponse)
	tocrawl := make(chan URL)
	go fetch(ts.URL, 1, resp, tocrawl)

	s := <-resp
	assert.Equal(ts.URL, s.Url)
	assert.Equal(`hogehoge<a href="http://example.com/">fugafuga`, s.Body)
	assert.Equal(200, s.StatusCode)

	tc := <-tocrawl
	assert.Equal(tc.Url, "http://example.com/")
	assert.Equal(tc.Depth, 0)
}

func TestCrawl(t *testing.T) {
	assert := assert.New(t)
	env.Init()
	err := models.Dialdb(env.GetDBHost())
	assert.Nil(err)
	defer models.Session.Close()

	ts := httptest.NewServer(http.HandlerFunc(DummyCrawledHandler))
	defer ts.Close()

	resp := make(chan CrawlResponse)
	go Crawl(ts.URL, 1, resp)

	s := <-resp
	assert.Equal(ts.URL, s.Url)
	assert.Equal(`hogehoge<a href="http://example.com/">fugafuga`, s.Body)
	assert.Equal(200, s.StatusCode)
}
