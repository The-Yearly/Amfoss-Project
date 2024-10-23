package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector()
	co := 0
	var s string
	var url string
	c.OnHTML("h3", func(h *colly.HTMLElement) {
		if co != 10 {
			fmt.Println(h.Text)
			co += 1
		} else {
			return
		}
	})
	c.OnHTML("a", func(h *colly.HTMLElement) {
		href := h.Attr("href")
		s = h.Text
		if s == "Next >" {
			nexturl := "https://www.google.com" + href
			c.Visit(nexturl)
		}
	})
	fmt.Println("Enter Search Word ")
	fmt.Scan(&s)
	sl := strings.ToLower(s)
	url = "https://www.google.com/search?q=" + sl
	c.Visit(url)
	c.Wait()
	fmt.Println(co)
}
