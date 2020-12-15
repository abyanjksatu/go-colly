package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main () {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla"),
	)

	c.OnHTML("a", func(e *colly.HTMLElement) {
		err := e.Request.Visit(e.Attr("href"))
		if err != nil {
			fmt.Println(err.Error())
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	err := c.Visit("https://www.zomato.com/id/jakarta/tebet-restoran")
	if err != nil {
		fmt.Println(err.Error())
	}
	
}