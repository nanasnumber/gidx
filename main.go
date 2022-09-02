package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"os"
	"regexp"
	"strings"
)

func clean(s string) string {
	r := strings.Replace(s, "/url?q=", "", -1)
	rg := regexp.MustCompile("(/[^&]+).*")
	result := rg.ReplaceAllString(r, `$1$2`)

	return result
}

func sourceList() []string {
	domain := os.Args[1]
	endpoint := "https://www.google.com/search?q=site%3A" + domain + "&oq=site%3A" + domain + "&num=1000&aqs=chrome..69i57j69i58.11371j0j4&sourceid=chrome&ie=UTF-8"
	scope := "body"
	node := "h3"
	var list []string

	c := colly.NewCollector()
	c.OnHTML(scope, func(e *colly.HTMLElement) {
		e.ForEach(node, func(idx int, item *colly.HTMLElement) {
			r, _ := item.DOM.ParentsUntil("~").Find("a").Attr("href")
			list = append(list, r)

		})
	})

	c.Visit(endpoint)

	return list
}

func indexedURL(s []string) []string {
	var r []string

	for _, url := range s {
		fmt.Println(clean(url))

		r = append(r, clean(url))
	}

	return r
}

func main() {
	indexedURL(sourceList())
}
