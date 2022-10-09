package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type item struct {
	Name   string `json:"name"`
	Price  string `json:"price"`
	ImgURL string `json:"imgurl"`
}

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("j2store.net"),
	)
	c.OnHTML("div[itemprop=itemListElement]", func(h *colly.HTMLElement) {
		fmt.Println(h.Text)

	})

	c.Visit("http://j2store.net/demo/index.php/shop")
}
