package crawler

import (
	"regexp"
)

func getAllLinks(body string) (urls []string)  {
	r := regexp.MustCompile(`https?://[\w/:%#\$&\?\(\)~\.=\+\-]+`)
	urls = r.FindAllString(body, -1)
	return urls
}
