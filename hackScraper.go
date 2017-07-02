package main

import (
	"fmt"
	"github.com/anaskhan96/soup"
)

func main() {

	resp, err := soup.Get("https://news.ycombinator.com")
	if err != nil {
		panic(err)
	}
	doc := soup.HTMLParse(resp)
	headline := doc.FindAll("a", "class", "storylink")
	links := doc.FindAll("span", "class", "sitestr")
	for i := 0; i <= len(headline); i++ {
		fmt.Println(headline[i].Text(), "=>", links[i].Text())
	}
}
