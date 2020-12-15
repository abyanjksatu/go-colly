package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main () {
	c := colly.NewCollector()

	c.OnHTML("a", func(e *colly.HTMLElement) {
		err := e.Request.Visit(e.Attr("href"))
		if err != nil {
			fmt.Println(err.Error())
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	err := c.Visit("http://go-colly.org/")
	if err != nil {
		fmt.Println(err.Error())
	}
	
}