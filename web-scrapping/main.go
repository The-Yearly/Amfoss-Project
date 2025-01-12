package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/gocolly/colly/v2"
)

var co = 0

func main() {
	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.83 Safari/537.36"
	c.Limit(&colly.LimitRule{
		Parallelism: 2,
		RandomDelay: 5,
	})
	var s string
	var url string
	fmt.Println("Enter Search Word")
	reader := bufio.NewReader(os.Stdin)
	s, _ = reader.ReadString('\n')
	sl := strings.TrimSpace(strings.ToLower(s))
	spaces := strings.Count(sl, " ")
	nsl := strings.Replace(sl, " ", "+", spaces)
	url = "https://www.google.com/search?q=" + nsl
	c.OnHTML("h3.LC20lb.MBeuO.DKV0Md", func(h *colly.HTMLElement) {
		if co != 10 {
			heading := h.Text
			p := h.DOM.Parent()
			link, _ := p.Attr("href")
			fmt.Println(co+1, ") Title: ", heading, "	Link: ", link, "\n")
			co += 1
		} else {
			return
		}
	})
	c.OnHTML("#pnnext", func(h *colly.HTMLElement) {
		if co < 10 {
			url = "https://www.google.com" + h.Attr("href")
			c.Visit(url)
		}
	})
	c.Visit(url)
	c.Wait()
	fmt.Println("Found ", co, "Results")
}
